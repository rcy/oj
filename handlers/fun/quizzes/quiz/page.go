package quiz

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/quizctx"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Router(r chi.Router) {
	r.Use(quizctx.Provider)
	r.Get("/", page)
	r.Post("/attempt", createAttempt)
}

func page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	l := layout.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := queries.QuizQuestions(ctx, quiz.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(questions) == 0 {
		render.Error(w, "quiz has no questions", http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout           layout.Data
		Quiz             api.Quiz
		Questions        []api.Question
		CreateAttemptURL string
	}{
		Layout:           l,
		Quiz:             quiz,
		Questions:        questions,
		CreateAttemptURL: r.URL.Path + "/attempt",
	})
}

func createAttempt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	queries := api.New(db.DB)
	attempt, err := queries.CreateAttempt(ctx, api.CreateAttemptParams{
		QuizID: quiz.ID,
		UserID: user.ID,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/fun/quizzes/attempts/%d", attempt.ID))
	w.WriteHeader(201)
}
