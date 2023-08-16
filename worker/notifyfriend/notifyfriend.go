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
		AID         int64     `db:"a_id"`
		BID         int64     `db:"b_id"`
		Email       string    `db:"email"`
		Username    string    `db:"username"`
		TargetEmail string    `db:"target_email"`
	}

	err = db.DB.Get(&friend, `
select
  f.id, f.created_at,
  a.id a_id, a.email, a.username,
  b.id b_id, b.email target_email
from friends f
join users a on a.id = f.a_id
join users b on b.id = f.b_id
where f.id = ?
`, j.Payload["id"])
	if err != nil {
		return err
	}

	var mutualID int64
	err = db.DB.Get(&mutualID, `select id from friends where a_id = ? and b_id = ?`, friend.BID, friend.AID)

	log.Printf("handleNotifyFriend %v, %d", friend, mutualID)

	var link url.URL
	var subject, emailBody string

	if mutualID != 0 {
		link = app.AbsoluteURL(url.URL{Path: fmt.Sprintf("/u/%d", friend.AID)})
		subject = fmt.Sprintf("%s accepted your friend request", friend.Username)
		emailBody = fmt.Sprintf("click here to view %s", link.String())
	} else {
		link = app.AbsoluteURL(url.URL{Path: "/connect"})
		subject = fmt.Sprintf("%s sent you a friend request", friend.Username)
		emailBody = fmt.Sprintf("click here to accept %s", link.String())
	}

	_, _, err = email.Send(subject, emailBody, friend.TargetEmail)

	return nil
}
