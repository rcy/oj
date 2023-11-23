package layout

import (
	_ "embed"
	"fmt"
	"html/template"
	"oj/api"
	"oj/db"
	"oj/element/gradient"
	"oj/services/background"
	"oj/templatehelpers"
)

var (
	//go:embed "layout.gohtml"
	layoutContent string
)

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

type Data struct {
	User               api.User
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func FromUser(user api.User) (Data, error) {
	backgroundGradient, err := background.ForUser(user.ID)
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
