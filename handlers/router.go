package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/chat"
	"oj/handlers/eventsource"
	"oj/handlers/header"
	"oj/handlers/parent"
	"oj/handlers/render"
	"oj/handlers/tools"
	"oj/handlers/u"
	auth "oj/handlers/welcome"
	"oj/models/users"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})
	r.Use(middleware.Logger)

	// authenticated routes
	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)

		r.Get("/", Home)

		r.Get("/header", header.Header)

		r.Get("/parent", parent.Index)
		r.Post("/parent/kids", parent.CreateKid)
		r.Delete("/parent/kids/{userID}", parent.DeleteKid)
		r.Post("/parent/kids/{userID}/logout", parent.LogoutKid)

		//r.Get("/chat/{roomID}", chat.Index)        // deprecated
		r.Post("/chat/messages", chat.PostChatMessage)
		r.Mount("/es", eventsource.SSE)

		r.Get("/tools", tools.Index)
		r.Post("/tools/picker", tools.Picker)
		r.Post("/tools/set-background", tools.SetBackground)

		r.Get("/u", u.UserIndex)
		r.Get("/u/{userID}", u.UserPage)
		r.Get("/u/{userID}/chat", chat.UserChatPage)

		r.Get("/bio", u.GetAbout)
		r.Get("/bio/edit", u.GetAboutEdit)
		r.Put("/bio", u.PutAbout)

		r.Get("/card/edit", u.GetCardEdit)
		r.Patch("/user", u.PatchUser)

		r.Get("/avatars", u.GetAvatars)
		r.Put("/avatar", u.PutAvatar)

		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			render.Error(w, "Page not found", 404)
		})
	})

	// non authenticated routes
	r.Route("/welcome", auth.Route)

	return r
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("kh_session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/welcome", http.StatusSeeOther)
				return
			} else {
				log.Printf("WARNING: weirderror: %s", err)
				http.Redirect(w, r, "/welcome?weirderror", http.StatusSeeOther)
				return
			}
		} else {
			var user users.User
			err := db.DB.Get(&user, "select users.* from sessions join users on sessions.user_id = users.id where sessions.key = ?", cookie.Value)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Redirect(w, r, "/welcome", http.StatusSeeOther)
					return
				} else {
					log.Printf("WARNING: weirderror2: %s", err)
					http.Redirect(w, r, "/welcome?weirderror2", http.StatusSeeOther)
					return
				}
			}
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
