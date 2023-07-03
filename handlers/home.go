package handlers

import (
	"fmt"
	"net/http"
	"oj/models/users"
)

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)

	if user.IsParent() {
		http.Redirect(w, r, "/parent", http.StatusFound)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/u/%d", user.ID), http.StatusFound)
	}
}
