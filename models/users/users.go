package users

import (
	"context"
	"database/sql"
	"oj/api"
	"oj/db"
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

// parents via kids_parents table
func GetParents(kidUserID int64) ([]User, error) {
	var parents []User

	err := db.DB.Select(&parents, "select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = ?", kidUserID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func CreateKid(u api.User, username string) (*User, error) {
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

func FindByStringId(stringID string) (api.User, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return api.User{}, err
	}
	queries := api.New(db.DB)
	return queries.UserByID(context.TODO(), int64(id))
}

func FindByUsername(username string) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select * from users where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindOrCreateParentByEmail(email string) (api.User, error) {
	ctx := context.TODO()
	queries := api.New(db.DB)

	// FIXME: make email not nullable and remove this
	nullableEmail := sql.NullString{String: email, Valid: true}

	user, err := queries.UserByEmail(ctx, nullableEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			// we don't have a username here, so use the email, they can change it later
			return queries.CreateParent(ctx, api.CreateParentParams{Email: nullableEmail, Username: email})
		}
		return api.User{}, err
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
