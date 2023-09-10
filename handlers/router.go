package handlers

import (
	"log"
	"net/http"
	"oj/handlers/chat"
	"oj/handlers/connect"
	"oj/handlers/connectkids"
	"oj/handlers/eventsource"
	"oj/handlers/family"
	"oj/handlers/friends"
	"oj/handlers/fun"
	"oj/handlers/header"
	"oj/handlers/me"
	"oj/handlers/me/editme"
	mw "oj/handlers/middleware"
	"oj/handlers/parent"
	"oj/handlers/render"
	"oj/handlers/tools"
	"oj/handlers/u"
	auth "oj/handlers/welcome"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})
	r.Use(middleware.Logger)

	// authenticated routes
	r.Route("/", func(r chi.Router) {
		r.Use(mw.Auth)
		r.Use(mw.Become)

		r.Get("/", Home)

		r.Get("/header", header.Header)

		r.Get("/parent", parent.Index)
		r.Post("/parent/kids", parent.CreateKid)
		r.Delete("/parent/kids/{userID}", parent.DeleteKid)
		r.Post("/parent/kids/{userID}/logout", parent.LogoutKid)

		r.Post("/chat/messages", chat.PostChatMessage)
		r.Mount("/es", eventsource.SSE)

		r.Get("/tools", tools.Index)
		r.Post("/tools/picker", tools.Picker)
		r.Post("/tools/set-background", tools.SetBackground)

		r.Get("/me", me.MyPage)
		r.Get("/me/edit", editme.MyPageEdit)
		r.Post("/me/edit", editme.Post)

		r.Get("/me/family", family.Page)
		r.Get("/me/friends", friends.Page)

		r.Get("/fun", fun.Page)

		r.Get("/u/{userID}", u.UserPage)
		r.Get("/u/{userID}/chat", chat.UserChatPage)

		r.Get("/avatars", editme.GetAvatars)
		r.Put("/avatar", editme.PutAvatar)

		r.Get("/connect", connect.Connect)
		r.Put("/connect/friend/{userID}", connect.PutParentFriend)
		r.Delete("/connect/friend/{userID}", connect.DeleteParentFriend)

		r.Get("/connectkids", connectkids.KidConnect)
		r.Put("/connectkids/friend/{userID}", connectkids.PutKidFriend)
		r.Delete("/connectkids/friend/{userID}", connectkids.DeleteKidFriend)

		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			render.Error(w, "Page not found", 404)
		})
	})

	// non authenticated routes
	r.Route("/welcome", auth.Route)

	// serve static files
	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fs))

	return r
}
