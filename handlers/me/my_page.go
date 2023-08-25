package me

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"
	"oj/util/hash"
)

var myPageTemplate = template.Must(template.New("layout.html").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/me/my_page.html"))

func MyPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ug, err := gradients.UserBackground(l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	canEdit := true

	d := struct {
		Layout  layout.Data
		User    users.User
		CanEdit bool
	}{
		Layout:  l,
		User:    l.User,
		CanEdit: canEdit,
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
