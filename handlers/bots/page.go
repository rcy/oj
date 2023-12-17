package bots

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/bots/ai"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
)

func strptr(str string) *string {
	return &str
}

var (
	//go:embed index.gohtml
	listPageContent  string
	listPageTemplate = layout.MustParse(listPageContent)

	//go:embed assistant.gohtml
	assistantPageContent  string
	assistantPageTemplate = layout.MustParse(assistantPageContent)

	//go:embed chat.gohtml
	chatPageContent  string
	chatPageTemplate = layout.MustParse(chatPageContent)
)

func Router(r chi.Router) {
	r.Get("/", listPage)
	r.Group(func(r chi.Router) {
		r.Use(provideAssistant)
		r.Get("/{assistantID}", assistantPage)
		r.Get("/{assistantID}/chat", chatRedirectPage)
		r.Get("/{assistantID}/chat/{threadID}", chatPage)
		r.Post("/{assistantID}/chat/{threadID}/messages", postMessage)
	})
}

func listPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	alist, err := ai.New().Client.ListAssistants(ctx, nil, nil, nil, nil)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, listPageTemplate, struct {
		Layout         layout.Data
		AssistantsList openai.AssistantsList
	}{
		Layout:         l,
		AssistantsList: alist,
	})
}

func assistantPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	assistant := assistantFromContext(ctx)

	render.Execute(w, assistantPageTemplate, struct {
		Layout    layout.Data
		Assistant openai.Assistant
	}{
		Layout:    l,
		Assistant: assistant,
	})
}

func chatRedirectPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := api.New(db.DB)
	client := ai.New().Client
	user := auth.FromContext(ctx)
	assistant := assistantFromContext(ctx)

	threads, err := query.AssistantThreads(ctx, api.AssistantThreadsParams{
		UserID:      user.ID,
		AssistantID: assistant.ID,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var threadID string
	if len(threads) > 0 {
		threadID = threads[0].ThreadID
	} else {
		thread, err := client.CreateThread(ctx, openai.ThreadRequest{})
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = query.CreateThread(ctx, api.CreateThreadParams{
			AssistantID: assistant.ID,
			ThreadID:    thread.ID,
			UserID:      user.ID,
		})
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		threadID = thread.ID
	}

	http.Redirect(w, r, r.URL.Path+"/"+threadID, http.StatusSeeOther)
}

func chatPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	l := layout.FromContext(ctx)
	client := ai.New().Client
	query := api.New(db.DB)
	assistant := assistantFromContext(ctx)

	userThread, err := query.UserThreadByID(ctx, api.UserThreadByIDParams{
		UserID:   user.ID,
		ThreadID: chi.URLParam(r, "threadID"),
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	thread, err := client.RetrieveThread(ctx, userThread.ThreadID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messagesList, err := client.ListMessage(ctx, thread.ID, nil, strptr("desc"), nil, nil)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, chatPageTemplate, struct {
		Layout    layout.Data
		Assistant openai.Assistant
		Thread    openai.Thread
		Messages  []openai.Message
		HasMore   bool
	}{
		Layout:    l,
		Assistant: assistant,
		Thread:    thread,
		Messages:  messagesList.Messages,
		HasMore:   messagesList.HasMore,
	})
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client := ai.New().Client
	assistant := assistantFromContext(ctx)

	content := strings.TrimSpace(r.FormValue("message"))
	if content == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadID")

	_, err := client.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = client.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: assistant.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, chatPageTemplate, "input", nil)
}
