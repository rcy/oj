package notifyfriend

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
	log.Printf("handleNotifyFriend job id: %d, payload: %v", j.ID, j.Payload)

	var friend struct {
		ID          int64
		CreatedAt   time.Time `db:"created_at"`
		Email       string    `db:"email"`
		Username    string    `db:"username"`
		TargetEmail string    `db:"target_email"`
	}

	err = db.DB.Get(&friend, `
select
  f.id, f.created_at,
  a.email, a.username,
  b.email target_email
from friends f
join users a on a.id = a_id
join users b on b.id = b_id
where f.id = ?
`, j.Payload["id"])
	if err != nil {
		return err
	}

	log.Printf("handleNotifyFriend %v", friend)

	link := app.AbsoluteURL(url.URL{Path: "/connect"})
	subject := fmt.Sprintf("%s sent you a friend request", friend.Username)
	emailBody := fmt.Sprintf("click here to accept %s", link.String())
	_, _, err = email.Send(subject, emailBody, friend.TargetEmail)

	return nil
}
