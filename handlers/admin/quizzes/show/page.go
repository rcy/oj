package show

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/middleware/quizctx"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(quizctx.Provider)
	r.Get("/", page)
	r.Patch("/", patchQuiz)
	r.Get("/edit", editQuiz)
	r.Post("/toggle-published", togglePublished)
	r.Get("/add-question", newQuestion)
	r.Post("/add-question", postNewQuestion)
	r.Get("/question/{questionID}/edit", editQuestion)
	r.Patch("/question/{questionID}", patchQuestion)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	l := layout.FromContext(ctx)
	quiz := quizctx.Value(ctx)

	questions, err := queries.QuizQuestions(ctx, quiz.ID)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout    layout.Data
		Quiz      api.Quiz
		Questions []api.Question
	}{
		Layout:    l,
		Quiz:      quiz,
		Questions: questions,
	})
}

func editQuiz(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())

	render.ExecuteNamed(w, pageTemplate, "quiz-header-edit", quiz)
}

func patchQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quiz := quizctx.Value(ctx)
	queries := api.New(db.DB)

	result, err := queries.UpdateQuiz(ctx, api.UpdateQuizParams{
		ID:          quiz.ID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "quiz-header", result)
}

func togglePublished(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())
	err := db.DB.Get(&quiz, `update quizzes set published = ? where id = ? returning *`, !quiz.Published.Bool, quiz.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, pageTemplate, "quiz-header", quiz)
}

func newQuestion(w http.ResponseWriter, r *http.Request) {
	quiz := quizctx.Value(r.Context())

	render.ExecuteNamed(w, pageTemplate, "new-question-form", struct{ QuizID int64 }{quiz.ID})
}

func editQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	questionID, _ := strconv.Atoi(chi.URLParam(r, "questionID"))

	quest, err := queries.Question(ctx, int64(questionID))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "edit-question-form", quest)
}

func postNewQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	quiz := quizctx.Value(r.Context())

	var err error
	var quest api.Question

	if r.FormValue("id") != "" {
		questionID, _ := strconv.Atoi(r.FormValue("id"))
		_, err = queries.Question(ctx, int64(questionID))
		if err != nil && err != sql.ErrNoRows {
			render.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		quest, err = queries.UpdateQuestion(r.Context(), api.UpdateQuestionParams{
			ID:     int64(questionID),
			Text:   r.FormValue("text"),
			Answer: r.FormValue("answer"),
		})
	} else {
		quest, err = queries.CreateQuestion(r.Context(), api.CreateQuestionParams{
			QuizID: quiz.ID,
			Text:   r.FormValue("text"),
			Answer: r.FormValue("answer"),
		})

	}

	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}

func patchQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	questionID, _ := strconv.Atoi(chi.URLParam(r, "questionID"))
	quest, err := queries.Question(ctx, int64(questionID))
	if err != nil {
		render.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	quest, err = queries.UpdateQuestion(r.Context(), api.UpdateQuestionParams{
		ID:     quest.ID,
		Text:   r.FormValue("text"),
		Answer: r.FormValue("answer"),
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}
