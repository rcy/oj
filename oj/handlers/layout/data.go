package layout

import (
	"net/http"
	"oj/db"
	"oj/element/gradient"
	"oj/models/gradients"
	"oj/models/users"
)

const File = "handlers/layout/layout.html"

type Data struct {
	User               users.User
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func GetData(r *http.Request) (Data, error) {
	user := users.Current(r)

	backgroundGradient, err := gradients.UserBackground(user.ID)
	if err != nil {
		return Data{}, err
	}

	var unreadCount int
	err = db.DB.Get(&unreadCount, `select count(*) from deliveries where recipient_id = ? and sent_at is null`, user.ID)
	if err != nil {
		return Data{}, err
	}

	return Data{
		User:               user,
		BackgroundGradient: *backgroundGradient,
		UnreadCount:        unreadCount,
	}, nil
}
