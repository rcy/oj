package layout

import (
	"net/http"
	"oj/element/gradient"
	"oj/models/gradients"
	"oj/models/users"
)

const File = "handlers/layout/layout.html"

type Data struct {
	User               users.User
	BackgroundGradient gradient.Gradient
}

func GetData(r *http.Request) (Data, error) {
	user := users.Current(r)

	backgroundGradient, err := gradients.UserBackground(user.ID)
	if err != nil {
		return Data{}, err
	}

	return Data{
		User:               user,
		BackgroundGradient: *backgroundGradient,
	}, nil
}
