package handlers

import (
	"fmt"
	"net/http"
	"oj/models/users"
)

func Home(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	if user.IsParent() {
		http.Redirect(w, r, "/parent", http.StatusFound)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/u/%s", user.Username), http.StatusFound)
	}
}
