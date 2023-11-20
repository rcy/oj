package completed

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/response"
	"strconv"

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

	quiz, err := queries.Quiz(ctx, attempt.QuizID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "attempt not found", http.StatusNotFound)
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

	responses, err := response.FindResponses(attempt.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout        layout.Data
		Quiz          api.Quiz
		Attempt       api.Attempt
		QuestionCount int64
		Responses     []response.Response
	}{
		Layout:        l,
		Quiz:          quiz,
		Attempt:       attempt,
		QuestionCount: questionCount,
		Responses:     responses,
	})
}
