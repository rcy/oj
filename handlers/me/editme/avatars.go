package editme

import (
	"fmt"
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/render"
	"oj/models/users"
	"oj/util/hash"
)

var avatarsTemplate = template.Must(template.New("avatars").ParseFiles("handlers/me/editme/avatars.gohtml"))

func GetAvatars(w http.ResponseWriter, r *http.Request) {
	const count = 27

	ctx := r.Context()
	user := users.FromContext(ctx)

	urls := []string{user.AvatarURL}

	for i := 0; i < count; i += 1 {
		url := fmt.Sprintf("https://robohash.org/%s?set=set5&size=80x80",
			hash.GenerateMD5(fmt.Sprintf("%d-%d", user.ID, i)))
		if url != urls[0] {
			urls = append(urls, url)
		}
	}

	render.ExecuteNamed(w, avatarsTemplate, "avatars", struct{ URLs []string }{urls})
}

func PutAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	newAvatarURL := r.FormValue("URL")

	err := db.DB.Get(&user, "update users set avatar_url = ? where id = ? returning *", newAvatarURL, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, avatarsTemplate, "changeable-avatar", struct{ User users.User }{user})
}
