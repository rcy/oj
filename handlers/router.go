package handlers

import (
	"log"
	"net/http"
	"oj/handlers/admin"
	"oj/handlers/chat"
	"oj/handlers/connect"
	"oj/handlers/connectkids"
	"oj/handlers/deliveries"
	"oj/handlers/eventsource"
	"oj/handlers/family"
	"oj/handlers/friends"
	"oj/handlers/fun"
	"oj/handlers/fun/chess"
	"oj/handlers/fun/gradients"
	"oj/handlers/fun/quizzes"
	"oj/handlers/fun/quizzes/attempt"
	"oj/handlers/fun/quizzes/attempt/completed"
	"oj/handlers/fun/quizzes/quiz"
	"oj/handlers/fun/stickers"
	"oj/handlers/header"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/me/editme"
	"oj/handlers/parent"
	"oj/handlers/render"
	"oj/handlers/u"
	"oj/handlers/welcome"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/become"
	"oj/internal/middleware/redirect"
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
		r.Use(auth.Provider)
		r.Use(become.Provider)
		r.Use(redirect.Redirect)
		r.Use(layout.Provider)
		r.Route("/", applicationRouter)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(auth.Provider)
		r.Use(layout.Provider)
		r.Route("/", admin.Router)
	})

	// non authenticated routes
	r.Route("/welcome", welcome.Route)

	// serve static files
	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/favicon.ico")
	})

	return r
}

func applicationRouter(r chi.Router) {
	r.Get("/", Home)

	r.Get("/header", header.Header)

	r.Get("/parent", parent.Index)
	r.Post("/parent/kids", parent.CreateKid)
	r.Delete("/parent/kids/{userID}", parent.DeleteKid)
	r.Post("/parent/kids/{userID}/logout", parent.LogoutKid)

	r.Post("/chat/messages", chat.PostChatMessage)
	r.Mount("/es", eventsource.SSE)

	r.Get("/me", me.Page)
	r.Get("/me/edit", editme.MyPageEdit)
	r.Post("/me/edit", editme.Post)

	r.Get("/me/family", family.Page)
	r.Get("/me/friends", friends.Page)

	r.Get("/fun", fun.Page)
	r.Get("/fun/gradients", gradients.Index)
	r.Post("/fun/gradients/picker", gradients.Picker)
	r.Post("/fun/gradients/set-background", gradients.SetBackground)

	r.Get("/fun/stickers", stickers.Page)
	r.Post("/fun/stickers", stickers.Submit)
	r.Post("/fun/stickers/save", stickers.SaveSticker)

	r.Get("/fun/chess", chess.Page)
	r.Get("/fun/chess/select/{rank}/{file}", chess.Select)
	r.Get("/fun/chess/unselect", chess.Unselect)
	//r.Get("/fun/chess/select/{r1}/{f1}/{r2}/{f2}", chess.Move)

	r.Get("/fun/quizzes", quizzes.Page)
	r.Route("/fun/quizzes/{quizID}", quiz.Router)
	r.Get("/fun/quizzes/attempts/{attemptID}", attempt.Page)
	r.Get("/fun/quizzes/attempts/{attemptID}/done", completed.Page)
	r.Post("/fun/quizzes/attempts/{attemptID}/question/{questionID}/response", attempt.PostResponse)

	r.Get("/u/{userID}", u.UserPage)
	r.Get("/u/{userID}/chat", chat.Page)

	r.Get("/avatars", editme.GetAvatars)
	r.Put("/avatar", editme.PutAvatar)

	r.Get("/connect", connect.Connect)
	r.Put("/connect/friend/{userID}", connect.PutParentFriend)
	r.Delete("/connect/friend/{userID}", connect.DeleteParentFriend)

	r.Get("/connectkids", connectkids.KidConnect)
	r.Put("/connectkids/friend/{userID}", connectkids.PutKidFriend)
	r.Delete("/connectkids/friend/{userID}", connectkids.DeleteKidFriend)

	r.Get("/deliveries/{deliveryID}", deliveries.Page)
	r.Get("/delivery/{deliveryID}", deliveries.Page) // temporary
	r.Post("/deliveries/{deliveryID}/logout", deliveries.Logout)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, "Page not found", 404)
	})
}
