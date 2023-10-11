package deliveries

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/deliveries"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed "page.gohtml"
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	l, err := layout.FromRequest(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deliveryID := chi.URLParam(r, "deliveryID")

	var delivery deliveries.Delivery

	query := `select * from deliveries where id = ?`

	err = db.DB.Get(&delivery, query, deliveryID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, err.Error(), http.StatusNotFound)
		} else {
			render.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if delivery.RecipientID == l.User.ID {
		url := fmt.Sprintf("/u/%d/chat", delivery.SenderID)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout          layout.Data
		LogoutActionURL string
		Delivery        deliveries.Delivery
	}{
		Layout:          l,
		LogoutActionURL: fmt.Sprintf("%d/logout", delivery.ID),
		Delivery:        delivery,
	})
}

// Logout and redirect back to delivery page to recheck current user
func Logout(w http.ResponseWriter, r *http.Request) {
	l, err := layout.FromRequest(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deliveryID := chi.URLParam(r, "deliveryID")

	var delivery deliveries.Delivery

	query := `select * from deliveries where id = ?`

	err = db.DB.Get(&delivery, query, deliveryID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, err.Error(), http.StatusNotFound)
		} else {
			render.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_, err = db.DB.Exec(`delete from sessions where user_id = ?`, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/deliveries/%d", delivery.ID)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
