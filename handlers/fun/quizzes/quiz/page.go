package quiz

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/attempts"
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

func Page(w http.ResponseWriter, r *http.Request) {
	l, err := layout.FromRequest(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quiz, err := quizzes.FindByStringID(chi.URLParam(r, "quizID"))
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "quiz not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		Quiz             *quizzes.Quiz
		Questions        []question.Question
		CreateAttemptURL string
	}{
		Layout:           l,
		Quiz:             quiz,
		Questions:        questions,
		CreateAttemptURL: r.URL.Path + "/attempt",
	})
}

func CreateAttempt(w http.ResponseWriter, r *http.Request) {
	user := users.FromContext(r.Context())

	quiz, err := quizzes.FindByStringID(chi.URLParam(r, "quizID"))
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "quiz not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	attempt, err := attempts.Create(quiz.ID, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Redirect", fmt.Sprintf("/fun/quizzes/attempts/%d", attempt.ID))
	w.WriteHeader(201)
}
