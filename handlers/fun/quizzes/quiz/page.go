package quiz

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/question"
	"oj/models/quizzes"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Router(r chi.Router) {
	r.Use(quizzes.Provider)
	r.Get("/", page)
	r.Post("/attempt", createAttempt)
}

func page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	quiz := quizzes.FromContext(r.Context())

	questions, err := quiz.FindQuestions()
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
		Quiz             quizzes.Quiz
		Questions        []question.Question
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
	quiz := quizzes.FromContext(ctx)

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
