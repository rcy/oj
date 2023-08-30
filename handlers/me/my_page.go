package me

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"oj/templatehelpers"
	"oj/util/hash"
	"sort"
)

var myPageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/me/my_page.gohtml"))

func MyPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type UserWithCount struct {
		users.User
		Role        string
		UnreadCount int
	}

	var friends []*UserWithCount
	err = db.DB.Select(&friends, `
select users.*, fi.b_role role
from users
join friends fi on fi.b_id = users.id and fi.a_id = $1
join friends fo on fo.a_id = users.id and fo.b_id = $1
where fi.b_role = 'friend'
`, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var family []*UserWithCount
	err = db.DB.Select(&family, `
select users.*, fi.b_role role
from users
join friends fi on fi.b_id = users.id and fi.a_id = $1
join friends fo on fo.a_id = users.id and fo.b_id = $1
where fi.b_role <> 'friend'
`, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Unread struct {
		SenderID int64 `db:"sender_id"`
		Count    int
	}
	var unreads []Unread

	err = db.DB.Select(&unreads, `
	  select sender_id, count(*) count
          from deliveries
          where recipient_id = ? and sent_at is null
          group by sender_id`, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, friend := range friends {
		for _, unread := range unreads {
			if unread.SenderID == friend.ID {
				friend.UnreadCount = unread.Count
				break
			}
		}
	}

	sort.Slice(friends, func(i, j int) bool {
		return friends[j].UnreadCount < friends[i].UnreadCount
	})

	d := struct {
		Layout  layout.Data
		User    users.User
		Friends []*UserWithCount
		Family  []*UserWithCount
	}{
		Layout:  l,
		User:    l.User,
		Friends: friends,
		Family:  family,
	}

	render.Execute(w, myPageTemplate, d)
}

func GetAboutEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	render.ExecuteNamed(w, myPageTemplate, "about-edit", struct{ User users.User }{User: user})
}

func PutAbout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	text := r.FormValue("text")

	err := db.DB.Get(&user, "update users set bio = ? where id = ? returning *", text, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, myPageTemplate, "about", struct {
		User    users.User
		CanEdit bool
	}{
		User:    user,
		CanEdit: true,
	})
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)

	render.ExecuteNamed(w, myPageTemplate, "about", struct {
		User    users.User
		CanEdit bool
	}{
		User:    user,
		CanEdit: true,
	})
}

func GetCardEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	render.ExecuteNamed(w, myPageTemplate, "card-edit", struct{ User users.User }{User: user})
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	username := r.FormValue("username")

	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("update users set username=? where id=?", username, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedUser, err := users.FindById(user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("user: %v %v", user, updatedUser)

	render.ExecuteNamed(w, myPageTemplate, "card", struct {
		User    users.User
		CanEdit bool
		Layout  layout.Data
	}{
		User:    *updatedUser,
		CanEdit: true,
		Layout:  l,
	})
}

func GetAvatars(w http.ResponseWriter, r *http.Request) {
	const count = 99

	ctx := r.Context()
	user := users.FromContext(ctx)

	urls := []string{user.AvatarURL}

	for i := 0; i < count; i += 1 {
		url := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro",
			hash.GenerateMD5(fmt.Sprintf("%s-%d", user.Username, i)))
		if url != urls[0] {
			urls = append(urls, url)
		}
	}

	render.ExecuteNamed(w, myPageTemplate, "avatars", struct{ URLs []string }{urls})
}

func PutAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	newAvatarURL := r.FormValue("URL")

	err := db.DB.Get(&user, "update users set avatar_url = ? where id = ? returning *", newAvatarURL, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, myPageTemplate, "changeable-avatar", struct{ User users.User }{user})
}
