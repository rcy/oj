package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"oj/db"
	"oj/handlers"
	"oj/handlers/auth"
	"oj/handlers/chat"
	"oj/handlers/parent"
	"oj/handlers/tools"
	"oj/handlers/u"

	"oj/models/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var err error
	var port string
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	err = db.DB.Ping()
	if err != nil {
		log.Fatal("could not ping db")
	}

	r := chi.NewRouter()

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})

	r.Use(middleware.Logger)
	//r.Use(authMiddleware)

	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/", handlers.Home)

		r.Route("/parent", parent.Route)
		r.Route("/chat", chat.Route)
		r.Route("/tools", tools.Route)

		r.Get("/u/{username}", u.UserPage)
	})

	r.Route("/welcome", auth.Route)

	http.Handle("/", r)

	log.Printf("listening on port %s", port)
	err = http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("server closed unexpectedly: %v\n", err)
		os.Exit(1)
	}
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
