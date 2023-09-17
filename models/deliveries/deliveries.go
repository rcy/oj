package deliveries

import (
	"time"
)

type Delivery struct {
	ID          int64
	CreatedAt   time.Time  `db:"created_at"`
	MessageID   int64      `db:"message_id"`
	RecipientID int64      `db:"recipient_id"`
	RoomID      int64      `db:"room_id"`
	SenderID    int64      `db:"sender_id"`
	SentAt      *time.Time `db:"sent_at"`
}
