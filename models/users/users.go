package users

import (
	"context"
	"database/sql"
	"errors"
	"html/template"
	"oj/db"
	"oj/md"
	"strconv"
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

func FromSessionKey(key string) (User, error) {
	var user User
	err := db.DB.Get(&user, "select users.* from sessions join users on sessions.user_id = users.id where sessions.key = ?", key)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

type contextKey int

const (
	userContextKey contextKey = iota
)

func NewContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func FromContext(ctx context.Context) User {
	return ctx.Value(userContextKey).(User)
}

var ErrNotAuthorized = errors.New("Not authorized")

func BecomeFromContext(ctx context.Context) (*User, error) {
	user := FromContext(ctx)
	if user.BecomeUserID == nil {
		return nil, nil
	}
	if !user.Admin {
		return nil, ErrNotAuthorized
	}
	return FindById(*user.BecomeUserID)
}

func (u User) BioHTML() template.HTML {
	return md.RenderString(u.Bio)
}

func (u User) Parents() ([]User, error) {
	return GetParents(u.ID)
}

// parents via kids_parents table
func GetParents(kidUserID int64) ([]User, error) {
	var parents []User

	err := db.DB.Select(&parents, "select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = ?", kidUserID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (u User) Kids() ([]User, error) {
	return GetKids(u.ID)
}

func (u User) CreateKid(username string) (*User, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec("insert into users(username) values(?)", username)
	if err != nil {
		return nil, err
	}

	kidID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	result, err = tx.Exec("insert into kids_parents(kid_id, parent_id) values(?, ?)", kidID, u.ID)
	if err != nil {
		return nil, err
	}

	result, err = tx.Exec("insert into friends(a_id, b_id, b_role) values(?, ?, 'child'),(?, ?, 'parent')", u.ID, kidID, kidID, u.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var kid User
	err = db.DB.Get(&kid, "select * from users where id = ?", kidID)
	if err != nil {
		return nil, err
	}

	return &kid, nil
}

func GetKids(parentUserID int64) ([]User, error) {
	var kids []User

	err := db.DB.Select(&kids, `
select users.* from kids_parents
join users on kids_parents.kid_id = users.id
where kids_parents.parent_id = ?
order by created_at desc
`, parentUserID)
	if err != nil {
		return nil, err
	}
	return kids, nil
}

func FindById(id int64) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select * from users where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindByStringId(stringID string) (*User, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return nil, err
	}
	return FindById(int64(id))
}

func FindByEmail(email string) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select * from users where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByUsername(username string) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select * from users where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateParent(email *string, username string) (*User, error) {
	result, err := db.DB.Exec("insert into users(email, username, is_parent) values(?, ?, true)", email, username)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return FindById(id)
}

func FindOrCreateParentByEmail(email string) (*User, error) {
	user, err := FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			// we don't have a username here, so use the email, they can change it later
			return CreateParent(&email, email)
		}
		return nil, err
	}
	return user, nil
}

func FindAll() ([]User, error) {
	var result []User
	err := db.DB.Select(&result, "select * from users order by created_at desc")
	if err != nil {
		return nil, err
	}
	return result, nil
}
