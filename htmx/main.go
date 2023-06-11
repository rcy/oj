package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"oj/handlers"
	"oj/handlers/auth"
	"oj/handlers/chat"
	"oj/handlers/games"
	"oj/handlers/tools"

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

	r := chi.NewRouter()

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})

	r.Use(middleware.Logger)
	//r.Use(authMiddleware)

	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/", handlers.Home)
		r.Route("/games", games.Route)
		r.Route("/chat", chat.Route)
		r.Route("/tools", tools.Route)
	})

	r.Route("/auth", auth.Route)

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
		cookie, err := r.Cookie("username")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/auth/signup?nocookie", 303)
			} else {
				http.Redirect(w, r, "/auth/signup?someothererror", 303)
			}
		} else {
			ctx := context.WithValue(r.Context(), "username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
