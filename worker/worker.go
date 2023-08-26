package worker

import (
	"context"
	"log"
	"oj/worker/notifydelivery"
	"oj/worker/notifyfriend"
	"time"

	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"github.com/acaloiaro/neoq/types"
)

var Queue types.Backend

func Start(ctx context.Context) error {
	var err error
	Queue, err = neoq.New(ctx)
	if err != nil {
		return err
	}

	Queue.Start(ctx, "notify-delivery", handler.New(notifydelivery.Handle))
	Queue.Start(ctx, "notify-friend", handler.New(notifyfriend.Handle))

	log.Print("started worker")

	return nil
}

func NotifyDelivery(deliveryID int64) (string, error) {
	return Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:    "notify-delivery",
		Payload:  map[string]any{"id": deliveryID},
		RunAfter: time.Now().Add(1 * time.Second),
	})
}

func NotifyFriend(friendID int64) (string, error) {
	log.Printf("Enqueue NotifyFriend %d", friendID)
	return Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:   "notify-friend",
		Payload: map[string]any{"id": friendID},
	})
}
