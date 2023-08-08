package worker

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"oj/app"
	"oj/db"
	"oj/services/email"
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

	Queue.Start(ctx, "notify-delivery", handler.New(handleNotifyDelivery))

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

func handleNotifyDelivery(ctx context.Context) error {
	j, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	log.Printf("handleNotifyDelivery job id: %d, payload: %v", j.ID, j.Payload)

	var delivery struct {
		ID             int64
		RecipientID    int64 `db:"recipient_id"`
		Username       string
		Email          *string
		SenderID       int64  `db:"sender_id"`
		SenderUsername string `db:"sender_username"`
		Body           string
		SentAt         *time.Time `db:"sent_at"`
	}

	err = db.DB.Get(&delivery, `
select
  d.id,
  r.username username,
  r.email email,
  r.id recipient_id,
  s.id sender_id,
  s.username sender_username,
  m.body body,
  sent_at sent_at
from deliveries d
join users r on r.id = d.recipient_id
join users s on s.id = d.sender_id
join messages m on m.id = d.message_id
where d.id = ?`, j.Payload["id"])
	if err != nil {
		return err
	}

	if delivery.Email == nil {
		return nil
	}

	if delivery.SenderID == delivery.RecipientID {
		return nil
	}

	if delivery.SentAt != nil {
		return nil
	}

	link := app.AbsoluteURL(url.URL{Path: fmt.Sprintf("/u/%d/chat", delivery.SenderID)})
	subject := fmt.Sprintf("%s sent you a message", delivery.SenderUsername)
	emailBody := fmt.Sprintf("%s %s", delivery.Body, link.String())
	_, _, err = email.Send(subject, emailBody, *delivery.Email)
	if err != nil {
		return err
	}

	return nil
}
