package attempt

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
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

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

	responseCount, err := attempt.ResponseCount()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextQuestion, err := attempt.NextQuestion()
	if err != nil {
		if err == sql.ErrNoRows {
			url := fmt.Sprintf("/fun/quizzes/attempts/%d/done", attempt.ID)
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Add("HX-Redirect", url)
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout        layout.Data
		Quiz          *quizzes.Quiz
		Attempt       *attempts.Attempt
		Question      *question.Question
		QuestionCount int
		ResponseCount int
	}{
		Layout:        l,
		Quiz:          quiz,
		Attempt:       attempt,
		Question:      nextQuestion,
		QuestionCount: questionCount,
		ResponseCount: responseCount,
	})
}

func PostResponse(w http.ResponseWriter, r *http.Request) {
	attemptID, err := strconv.Atoi(chi.URLParam(r, "attemptID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	questionID, err := strconv.Atoi(chi.URLParam(r, "questionID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	text := strings.TrimSpace(r.FormValue("response"))
	if text != "" {
		_, err := attempts.CreateResponse(int64(attemptID), int64(questionID), text)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("HX-Location", fmt.Sprintf("/fun/quizzes/attempts/%d", attemptID))
	w.WriteHeader(201)
}
