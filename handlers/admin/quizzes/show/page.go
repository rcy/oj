package show

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/question"
	"oj/models/quizzes"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(quizzes.Provider)
	r.Get("/", page)
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
	l := layout.FromContext(r.Context())

	quiz := quizzes.FromContext(r.Context())

	questions, err := quiz.FindQuestions()
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout    layout.Data
		Quiz      quizzes.Quiz
		Questions []question.Question
	}{
		Layout:    l,
		Quiz:      quiz,
		Questions: questions,
	})
}

func togglePublished(w http.ResponseWriter, r *http.Request) {
	quiz := quizzes.FromContext(r.Context())
	err := db.DB.Get(&quiz, `update quizzes set published = ? where id = ? returning *`, !quiz.Published, quiz.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.ExecuteNamed(w, pageTemplate, "quiz-header", quiz)
}

func newQuestion(w http.ResponseWriter, r *http.Request) {
	quiz := quizzes.FromContext(r.Context())

	render.ExecuteNamed(w, pageTemplate, "new-question-form", struct{ QuizID int64 }{quiz.ID})
}

func editQuestion(w http.ResponseWriter, r *http.Request) {
	quest, err := question.FindByStringID(chi.URLParam(r, "questionID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "edit-question-form", quest)
}

func postNewQuestion(w http.ResponseWriter, r *http.Request) {
	quiz := quizzes.FromContext(r.Context())

	var err error
	var quest *question.Question

	if r.FormValue("id") != "" {
		quest, err = question.FindByStringID(r.FormValue("id"))
		if err != nil && err != sql.ErrNoRows {
			render.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	} else {
		quest = &question.Question{}
	}

	quest.QuizID = quiz.ID
	quest.Text = r.FormValue("text")
	quest.Answer = r.FormValue("answer")

	quest, err = quest.Save(r.Context(), db.DB)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}

func patchQuestion(w http.ResponseWriter, r *http.Request) {
	quest, err := question.FindByStringID(chi.URLParam(r, "questionID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	quest.Text = r.FormValue("text")
	quest.Answer = r.FormValue("answer")

	quest, err = quest.Save(r.Context(), db.DB)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "question", quest)
}
