package notifydelivery

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"oj/app"
	"oj/db"
	"oj/services/email"
	"time"

	"github.com/acaloiaro/neoq/jobs"
)

func Handle(ctx context.Context) error {
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

	link := app.AbsoluteURL(url.URL{Path: fmt.Sprintf("/delivery/%d", delivery.ID)})
	subject := fmt.Sprintf("%s sent you a message", delivery.SenderUsername)
	emailBody := fmt.Sprintf("%s\n\nclick here to reply: %s", delivery.Body, link.String())
	_, _, err = email.Send(subject, emailBody, *delivery.Email)
	if err != nil {
		return err
	}

	return nil
}
