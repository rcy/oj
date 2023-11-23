package auth

import (
	"context"
	"errors"
	"net/http"
	"oj/api"
	"oj/db"
	"time"
)

// Provide user in context based on session cookie, redirecting to login if no session found
func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		queries := api.New(db.DB)

		cookie, err := r.Cookie("kh_session")
		if err != nil {
			redirectToLogin(w, r)
			return
		}

		user, err := queries.UserBySessionKey(ctx, cookie.Value)
		if err != nil {
			redirectToLogin(w, r)
			return
		}

		ctx = NewContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey int

const (
	userContextKey contextKey = iota
)

func NewContext(ctx context.Context, user api.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func FromContext(ctx context.Context) api.User {
	return ctx.Value(userContextKey).(api.User)
}

var ErrNotAuthorized = errors.New("Not authorized")

// Save the current path in a cookie and redirect to welcome page
func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "redirect",
		Value:   r.URL.Path,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour)})

	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
