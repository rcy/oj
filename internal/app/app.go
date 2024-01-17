package app

import (
	"net/http"
	"oj/handlers/bots"
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
	"oj/handlers/humans"
	"oj/handlers/me"
	"oj/handlers/me/editme"
	"oj/handlers/parent"
	"oj/handlers/postoffice"
	"oj/handlers/u"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type handler struct {
	router   *chi.Mux
	database *sqlx.DB
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func Handler(database *sqlx.DB) *handler {
	h := &handler{router: chi.NewRouter(), database: database}

	h.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/me", http.StatusFound)
	})

	h.router.Get("/header", header.Header)

	h.router.Get("/parent", parent.Index)
	h.router.Post("/parent/kids", parent.CreateKid)
	h.router.Delete("/parent/kids/{userID}", parent.DeleteKid)
	h.router.Post("/parent/kids/{userID}/logout", parent.LogoutKid)

	h.router.Post("/chat/messages", chat.PostChatMessage)
	h.router.Mount("/es", eventsource.SSE)

	h.router.Get("/me", me.Page)
	h.router.Get("/me/edit", editme.MyPageEdit)
	h.router.Post("/me/edit", editme.Post)
	h.router.Get("/avatars", editme.GetAvatars)
	h.router.Put("/avatar", editme.PutAvatar)

	h.router.Get("/me/humans", humans.Page)
	h.router.Get("/me/family", family.Page)
	h.router.Get("/me/friends", friends.Page)

	h.router.Get("/fun", fun.Page)
	h.router.Get("/fun/gradients", gradients.Index)
	h.router.Post("/fun/gradients/picker", gradients.Picker)
	h.router.Post("/fun/gradients/set-background", gradients.SetBackground)

	h.router.Get("/fun/stickers", stickers.Page)
	h.router.Post("/fun/stickers", stickers.Submit)
	h.router.Post("/fun/stickers/save", stickers.SaveSticker)

	h.router.Get("/fun/chess", chess.Page)
	h.router.Get("/fun/chess/select/{rank}/{file}", chess.Select)
	h.router.Get("/fun/chess/unselect", chess.Unselect)
	//h.router.Get("/fun/chess/select/{r1}/{f1}/{r2}/{f2}", chess.Move)

	h.router.Get("/fun/quizzes", quizzes.Page)
	h.router.Route("/fun/quizzes/{quizID}", quiz.Router)
	h.router.Get("/fun/quizzes/attempts/{attemptID}", attempt.Page)
	h.router.Get("/fun/quizzes/attempts/{attemptID}/done", completed.Page)
	h.router.Post("/fun/quizzes/attempts/{attemptID}/question/{questionID}/response", attempt.PostResponse)

	h.router.Route("/bots", bots.Router)

	h.router.Route("/u/{userID}", u.Router)
	h.router.Get("/u/{userID}/chat", chat.Page)

	h.router.Get("/connect", connect.Connect)
	h.router.Put("/connect/friend/{userID}", connect.PutParentFriend)
	h.router.Delete("/connect/friend/{userID}", connect.DeleteParentFriend)

	h.router.Get("/connectkids", connectkids.KidConnect)
	h.router.Put("/connectkids/friend/{userID}", connectkids.PutKidFriend)
	h.router.Delete("/connectkids/friend/{userID}", connectkids.DeleteKidFriend)

	h.router.Get("/deliveries/{deliveryID}", deliveries.Page)
	h.router.Get("/delivery/{deliveryID}", deliveries.Page) // temporary
	h.router.Post("/deliveries/{deliveryID}/logout", deliveries.Logout)

	h.router.Route("/postoffice", postoffice.Router)

	return h
}
