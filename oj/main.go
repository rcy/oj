package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"oj/db"
	"oj/handlers"
)

func main() {
	err := db.DB.Ping()
	if err != nil {
		log.Fatalf("could not ping db: %s", err)
	}

	listenAndServe(os.Getenv("PORT"), handlers.Router())
}

func listenAndServe(port string, handler http.Handler) {
	if port == "" {
		port = "8080"
	}

	http.Handle("/", handler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("server closed unexpectedly: %v\n", err)
		os.Exit(1)
	}
}
