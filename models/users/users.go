package users

import (
	"database/sql"
	"net/http"
	"oj/db"
	"strconv"
	"time"
)

type User struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	Username  string
	Email     *string
	AvatarURL string `db:"avatar_url"`
}

func (u User) IsParent() bool {
	return u.Email != nil
}

func (u User) Parents() ([]User, error) {
	return GetParents(u.ID)
}

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

func Current(r *http.Request) User {
	return r.Context().Value("user").(User)
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

func Create(email *string, username string) (*User, error) {
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
			return Create(&email, email)
		}
		return nil, err
	}
	return user, nil
}

func FindAll() ([]User, error) {
	var result []User
	err := db.DB.Select(&result, "select * from users")
	if err != nil {
		return nil, err
	}
	return result, nil
}
