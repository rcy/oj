package handlers

import (
	"net/http"
	"time"
)

func GetSignout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "username", Expires: time.Now()})
	http.Redirect(w, r, "/", http.StatusFound)
}
