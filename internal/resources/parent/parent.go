package parent

import (
	"database/sql"
	_ "embed"
	"errors"
	"log"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/services/family"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	DB    *sqlx.DB
	Model *api.Queries
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.index)
	r.Post("/kids", rs.createKid)
	r.Delete("/kids/{userID}", rs.deleteKid)
	r.Post("/kids/{userID}/logout", rs.logoutKid)
	return r
}

var (
	//go:embed parent.gohtml
	pageContent string

	t = layout.MustParse(pageContent)
)

func (rs Resource) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())

	kids, err := rs.Model.KidsByParentID(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	err = t.Execute(w, struct {
		Layout layout.Data
		User   api.User
		Kids   []api.User
	}{
		Layout: l,
		User:   l.User,
		Kids:   kids,
	})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func (rs Resource) createKid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	err := r.ParseForm()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	username := r.PostForm.Get("username")

	_, err = rs.Model.UserByUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		kid, err := family.CreateKid(ctx, user.ID, username)
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}
		log.Printf("kid: %v", kid)
		http.Redirect(w, r, "/parent", http.StatusSeeOther)
		return
	}

	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	render.Error(w, "username taken", http.StatusConflict)
	return
}

func (rs Resource) deleteKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	tx, err := rs.DB.Beginx()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
delete from kids_parents where kid_id = ?;
delete from bios where user_id = ?;
delete from deliveries where sender_id = ? or recipient_id = ?;
delete from gradients where user_id = ?;
delete from kids_codes where user_id = ?;
delete from messages where sender_id = ?;
delete from room_users where user_id = ?;
delete from sessions where user_id = ?;
delete from users where id = ?;
`, userID, userID, userID, userID, userID, userID, userID, userID, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs Resource) logoutKid(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	_, err := rs.DB.Exec(`delete from sessions where user_id = ?`, userID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
