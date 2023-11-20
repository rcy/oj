package attempt

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/quizzes"
	"oj/models/users"
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
	ctx := r.Context()
	l := layout.FromContext(ctx)
	queries := api.New(db.DB)

	attemptID, _ := strconv.Atoi(chi.URLParam(r, "attemptID"))
	attempt, err := queries.GetAttemptByID(ctx, int64(attemptID))
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
			render.Error(w, "quiz not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	questionCount, err := queries.QuestionCount(ctx, quiz.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseCount, err := queries.ResponseCount(ctx, attempt.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nextQuestion, err := queries.AttemptNextQuestion(ctx, api.AttemptNextQuestionParams{
		QuizID:    attempt.QuizID,
		AttemptID: attempt.ID,
	})
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
		Attempt       api.Attempt
		Question      api.Question
		QuestionCount int64
		ResponseCount int64
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
	ctx := r.Context()
	queries := api.New(db.DB)

	user := users.FromContext(ctx)

	attemptID, err := strconv.Atoi(chi.URLParam(r, "attemptID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	attempt, err := queries.GetAttemptByID(ctx, int64(attemptID))
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
		_, err := queries.CreateResponse(ctx, api.CreateResponseParams{
			QuizID:     attempt.QuizID,
			UserID:     user.ID,
			AttemptID:  attemptID,
			QuestionID: questionID,
			Text:       text,
		})

		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("HX-Location", fmt.Sprintf("/fun/quizzes/attempts/%d", attemptID))
	w.WriteHeader(201)
}
