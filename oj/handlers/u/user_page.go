package u

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"
	"time"

	"github.com/go-chi/chi/v5"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/u/user_page.html"))

func UserPage(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	username := chi.URLParam(r, "username")
	user, err := users.FindByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", 404)
			return
		}
	}
	ug, err := gradients.UserBackground(user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	bio, err := getBio(user.ID)
	if err != nil {
		log.Print("err(R2Hd): ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    users.User
		Bio     Bio
		CanEdit bool
	}{
		Layout:  l,
		User:    *user,
		Bio:     *bio,
		CanEdit: false,
	}

	render.Execute(w, t, d)
}

func MePage(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	bio, err := getBio(l.User.ID)
	if err != nil {
		log.Print("err(R2Hd): ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    users.User
		Bio     Bio
		CanEdit bool
	}{
		Layout:  l,
		User:    l.User,
		Bio:     *bio,
		CanEdit: true,
	}

	render.Execute(w, t, d)
}

func GetAboutEdit(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	bio, err := getBio(user.ID)
	if err != nil {
		log.Print("err(a4gt): ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "about-edit", struct{ Bio Bio }{Bio: *bio})
}

func PutAbout(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "empty text", http.StatusBadRequest)
		return
	}
	var bio Bio
	err := db.DB.Get(&bio, "insert into bios(user_id,text) values(?,?) returning *", user.ID, text)
	if err != nil {
		log.Print("error(owcE): ", err)
		http.Error(w, "unknown error", http.StatusInternalServerError)
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
		log.Print("err(f8Ns): ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
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
