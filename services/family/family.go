package family

import (
	"context"
	"database/sql"
	"oj/api"
	"oj/db"
)

func CreateKid(ctx context.Context, u api.User, username string) (*api.User, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	queries := api.New(tx)

	kid, err := queries.CreateUser(ctx, username)
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateKidParent(ctx, api.CreateKidParentParams{KidID: kid.ID, ParentID: u.ID})
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateFriend(ctx, api.CreateFriendParams{AID: u.ID, BID: kid.ID, BRole: "child"})
	if err != nil {
		return nil, err
	}

	_, err = queries.CreateFriend(ctx, api.CreateFriendParams{AID: kid.ID, BID: u.ID, BRole: "parent"})
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &kid, nil
}

func FindOrCreateParentByEmail(ctx context.Context, email string) (api.User, error) {
	queries := api.New(db.DB)

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
