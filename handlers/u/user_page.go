package u

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/md"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"
	"oj/util/hash"
	"time"

	"github.com/go-chi/chi/v5"
)

var t = template.Must(template.New("layout.html").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/u/user_page.html"))

func UserPage(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	user, err := users.FindByStringId(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ug, err := gradients.UserBackground(user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	bio, err := getBio(user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	canEdit := l.User.ID == user.ID

	d := struct {
		Layout  layout.Data
		User    users.User
		Bio     Bio
		CanEdit bool
	}{
		Layout:  l,
		User:    *user,
		Bio:     *bio,
		CanEdit: canEdit,
	}

	render.Execute(w, t, d)
}

func GetAboutEdit(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	bio, err := getBio(user.ID)
	if err != nil {
		render.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "about-edit", struct{ Bio Bio }{Bio: *bio})
}

func PutAbout(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	text := r.FormValue("text")
	if text == "" {
		render.Error(w, "empty text", http.StatusBadRequest)
		return
	}
	var bio Bio
	err := db.DB.Get(&bio, "insert into bios(user_id,text) values(?,?) returning *", user.ID, text)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "about", struct {
		Bio     Bio
		CanEdit bool
	}{
		Bio:     bio,
		CanEdit: true,
	})
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	bio, err := getBio(user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "about", struct {
		Bio     Bio
		CanEdit bool
	}{
		Bio:     *bio,
		CanEdit: true,
	})
}

type Bio struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	UserID    int64     `db:"user_id"`
	Text      string
}

func (b Bio) HTML() template.HTML {
	return md.RenderString(b.Text)
}

func getBio(userID int64) (*Bio, error) {
	var bio Bio
	err := db.DB.Get(&bio, "select * from bios where user_id = ? order by created_at desc limit 1", userID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	return &bio, nil
}

func GetCardEdit(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	render.ExecuteNamed(w, t, "card-edit", struct{ User users.User }{User: user})
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	username := r.FormValue("username")

	l, err := layout.GetData(r)
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

	render.ExecuteNamed(w, t, "card", struct {
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

	user := users.Current(r)

	urls := []string{user.AvatarURL}

	for i := 0; i < count; i += 1 {
		url := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro", hash.GenerateMD5(fmt.Sprintf("%s-%d", user.Username, i)))
		if url != urls[0] {
			urls = append(urls, url)
		}
	}

	render.ExecuteNamed(w, t, "avatars", struct{ URLs []string }{urls})
}

func PutAvatar(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	newAvatarURL := r.FormValue("URL")

	err := db.DB.Get(&user, "update users set avatar_url = ? where id = ? returning *", newAvatarURL, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, t, "changeable-avatar", struct{ User users.User }{user})
}
