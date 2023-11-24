package users

import (
	"time"
)

type User struct {
	ID           int64
	CreatedAt    time.Time `db:"created_at"`
	Username     string
	Email        *string
	AvatarURL    string `db:"avatar_url"`
	IsParent     bool   `db:"is_parent"`
	Bio          string `db:"bio"`
	BecomeUserID *int64 `db:"become_user_id"`
	Admin        bool
}
