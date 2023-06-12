package users

import (
	"database/sql"
	"net/http"
	"oj/db"
	"time"
)

type User struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	Username  string
	Email     *string
}

func Current(r *http.Request) User {
	return r.Context().Value("user").(User)
}

func FindById(id int64) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select id, created_at, username, email from users where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindByEmail(email string) (*User, error) {
	var user User

	err := db.DB.Get(&user, "select id, created_at, username, email from users where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(email string, username string) (*User, error) {
	result, err := db.DB.Exec("insert into users(email, username) values(?, ?)", email, username)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return FindById(id)
}

func FindOrCreateByEmail(email string) (*User, error) {
	user, err := FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			// we don't have a username here, so use the email, they can change it later
			return Create(email, email)
		}
		return nil, err
	}
	return user, nil
}
