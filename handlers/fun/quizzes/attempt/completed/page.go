package completed

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/attempts"
	"oj/models/quizzes"
	"oj/models/response"

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

	attempt, err := attempts.FindByStringID(chi.URLParam(r, "attemptID"))
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "attempt not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quiz, err := quizzes.FindByID(attempt.QuizID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "attempt not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	questionCount, err := attempt.QuestionCount()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses, err := response.FindResponses(attempt.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout        layout.Data
		Quiz          *quizzes.Quiz
		Attempt       *attempts.Attempt
		QuestionCount int
		Responses     []response.Response
	}{
		Layout:        l,
		Quiz:          quiz,
		Attempt:       attempt,
		QuestionCount: questionCount,
		Responses:     responses,
	})
}
