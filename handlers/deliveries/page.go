package deliveries

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed "page.gohtml"
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())
	queries := api.New(db.DB)

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := queries.Delivery(ctx, int64(deliveryID))
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
		Delivery        api.Delivery
	}{
		Layout:          l,
		LogoutActionURL: fmt.Sprintf("%d/logout", delivery.ID),
		Delivery:        delivery,
	})
}

// Logout and redirect back to delivery page to recheck current user
func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())
	queries := api.New(db.DB)

	deliveryID, _ := strconv.Atoi(chi.URLParam(r, "deliveryID"))
	delivery, err := queries.Delivery(ctx, int64(deliveryID))
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
