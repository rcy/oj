package worker

import (
	"context"
	"log"
	"oj/worker/notifydelivery"
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

	log.Print("started worker")

	return nil
}

func NotifyDelivery(deliveryID int64) {
	Queue.Enqueue(context.Background(), &jobs.Job{
		Queue:    "notify-delivery",
		Payload:  map[string]any{"id": deliveryID},
		RunAfter: time.Now().Add(1 * time.Second),
	})
}
