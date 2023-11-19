package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"oj/db"
	"oj/handlers"
	"oj/handlers/eventsource"
	"oj/services/email"
	"oj/worker"

	"github.com/alexandrevicenzi/go-sse"
)

func main() {
	err := db.DB.Ping()
	if err != nil {
		log.Fatalf("could not ping db: %s", err)
	}

	err = worker.Start(context.Background())
	if err != nil {
		log.Fatalf("could not start worker: %s", err)
	}

	go func() {
		count := 0
		for {
			id := fmt.Sprint(count)
			data := time.Now().Format(time.RFC3339Nano)
			eventsource.SSE.SendMessage("", sse.NewMessage(id, data, "KEEP_ALIVE"))
			count += 1
			time.Sleep(30 * time.Second)
		}
	}()

	err = email.Send("oj startup", "application started", os.Getenv("DEV_EMAIL"))
	if err != nil {
		log.Fatalf("failed to send startup email: %s", err)
	}

	listenAndServe(os.Getenv("PORT"), handlers.Router())
}

func listenAndServe(port string, handler http.Handler) {
	if port == "" {
		port = "8080"
	}

	http.Handle("/", handler)

	log.Printf("listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("server closed unexpectedly: %v\n", err)
		os.Exit(1)
	}
}
