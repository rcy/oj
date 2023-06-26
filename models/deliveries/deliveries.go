package deliveries

import (
	"time"
)

type Delivery struct {
	ID          int64
	CreatedAt   time.Time `db:"created_at"`
	MessageID   int64     `db:"message_id"`
	RecipientID int64     `db:"recipient_id"`
	SentAt      time.Time `db:"sent_at"`
	ReadAt      time.Time `db:"read_at"`
}
