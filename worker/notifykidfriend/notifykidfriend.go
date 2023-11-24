package notifykidfriend

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"oj/api"
	"oj/app"
	"oj/db"
	"oj/services/email"
	"time"

	"github.com/acaloiaro/neoq/jobs"
)

func Handle(ctx context.Context) error {
	queries := api.New(db.DB)

	j, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	log.Printf("handleNotifyKidFriend job id: %d, payload: %v", j.ID, j.Payload)

	var friend struct {
		ID        int64
		CreatedAt time.Time `db:"created_at"`
		AID       int64     `db:"a_id"`
		BID       int64     `db:"b_id"`
		AUsername string    `db:"a_username"`
		BUsername string    `db:"b_username"`
	}

	err = db.DB.Get(&friend, `
select
  f.id, f.created_at,
  a.id a_id,
  a.username a_username,
  b.id b_id,
  b.username b_username
from friends f
join users a on a.id = f.a_id
join users b on b.id = f.b_id
where f.id = ?
`, j.Payload["id"])
	if err != nil {
		return fmt.Errorf("getting friend %w", err)
	}

	var mutualID int64
	err = db.DB.Get(&mutualID, `select id from friends where a_id = ? and b_id = ?`, friend.BID, friend.AID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("getting mutual %w", err)
	}

	aUserLink := app.AbsoluteURL(url.URL{Path: fmt.Sprintf("/u/%d", friend.AID)})

	bParents, err := queries.ParentsByKidID(ctx, friend.BID)
	if err != nil {
		return fmt.Errorf("GetParents %w", err)
	}

	for _, bParent := range bParents {
		var subject, emailBody string

		if mutualID != 0 {
			subject = fmt.Sprintf("%s accepted a friend request from your child, %s", friend.AUsername, friend.BUsername)
			emailBody = fmt.Sprintf("click here to view %s: %s", friend.AUsername, aUserLink.String())
		} else {
			subject = fmt.Sprintf("%s sent a friend request to your child, %s", friend.AUsername, friend.BUsername)
			emailBody = fmt.Sprintf("click here to view %s: %s", friend.AUsername, aUserLink.String())
		}
		err = email.Send(subject, emailBody, bParent.Email.String)
		if err != nil {
			return err
		}
	}

	aParents, err := queries.ParentsByKidID(ctx, friend.AID)
	if err != nil {
		return err
	}
	for _, aParent := range aParents {
		var subject, emailBody string

		if mutualID != 0 {
			subject = fmt.Sprintf("your child, %s, accepted a friend request from %s", friend.AUsername, friend.BUsername)
			emailBody = fmt.Sprintf("click here to view %s: %s", friend.BUsername, aUserLink.String())
		} else {
			subject = fmt.Sprintf("your child, %s, sent a friend request to %s", friend.BUsername, friend.AUsername)
			emailBody = fmt.Sprintf("click here to view %s: %s", friend.BUsername, aUserLink.String())
		}
		err = email.Send(subject, emailBody, aParent.Email.String)
		if err != nil {
			return err
		}
	}

	return nil
}
