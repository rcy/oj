package layout

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"oj/db"
	"oj/element/gradient"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"
)

//go:embed "layout.gohtml"
var layoutContent string

func MustParse(templateContent ...string) *template.Template {
	tpl := template.New("layout").Funcs(templatehelpers.FuncMap)

	for i, content := range append([]string{layoutContent}, templateContent...) {
		var err error
		tpl, err = tpl.Parse(content)
		if err != nil {
			fmt.Println(i, content)
			panic(err)
		}
	}
	return tpl
}

const File = "handlers/layout/layout.gohtml"

type Data struct {
	User               users.User
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func FromContext(ctx context.Context) (Data, error) {
	user := users.FromContext(ctx)

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
