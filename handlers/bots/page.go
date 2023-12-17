package bots

import (
	_ "embed"
	"fmt"
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

	//go:embed create.gohtml
	createPageContent  string
	createPageTemplate = layout.MustParse(createPageContent)

	//go:embed chat.gohtml
	chatPageContent  string
	chatPageTemplate = layout.MustParse(chatPageContent)
)

func Router(r chi.Router) {
	r.Get("/", listPage)
	r.Get("/create", createPage)
	r.Post("/create", postCreate)
	r.Group(func(r chi.Router) {
		r.Use(provideAssistant)
		r.Get("/{assistantID}", assistantPage)
		r.Get("/{assistantID}/chat", chatRedirectPage)
		r.Get("/{assistantID}/chat/{threadID}", chatPage)
		r.Post("/{assistantID}/chat/{threadID}/messages", postMessage)
		r.Get("/{assistantID}/chat/{threadID}/runstatus/{runID}", getRunStatus)
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

func createPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	render.Execute(w, createPageTemplate, struct {
		Layout layout.Data
	}{
		Layout: l,
	})
}

func postCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client := ai.New().Client

	name := r.FormValue("name")
	if name == "" {
		http.Redirect(w, r, "/bots/create", http.StatusSeeOther)
		return
	}
	instructions := fmt.Sprintf("Your name is %s. %s", name, r.FormValue("instructions"))

	// models, err := client.ListModels(ctx)
	// if err != nil {
	// 	render.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// for _, m := range models.Models {
	// 	fmt.Printf("%v\n", m)
	// }

	model := "gpt-3.5-turbo"

	asst, err := client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        model,
		Name:         &name,
		Instructions: &instructions,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/bots/"+asst.ID, http.StatusSeeOther)
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

	run, err := client.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: assistant.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// trigger update to show user's message in the chat
	w.Header().Add("HX-Trigger", "messagesUpdated")

	render.ExecuteNamed(w, chatPageTemplate, "thinking", struct {
		Assistant openai.Assistant
		Run       openai.Run
	}{
		Assistant: assistant,
		Run:       run,
	})
}

func getRunStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client := ai.New().Client
	assistant := assistantFromContext(ctx)

	run, err := client.RetrieveRun(ctx, chi.URLParam(r, "threadID"), chi.URLParam(r, "runID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch run.Status {
	case openai.RunStatusQueued, openai.RunStatusInProgress:
		render.ExecuteNamed(w, chatPageTemplate, "thinking", struct {
			Assistant openai.Assistant
			Run       openai.Run
		}{
			Assistant: assistant,
			Run:       run,
		})
	default:
		// the run may or may not have been successful, but at this point, we want to
		// trigger an event to update the chat messages
		w.Header().Add("HX-Trigger", "messagesUpdated")

		thread, err := client.RetrieveThread(ctx, run.ThreadID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.ExecuteNamed(w, chatPageTemplate, "input", struct {
			Assistant openai.Assistant
			Thread    openai.Thread
		}{
			Assistant: assistant,
			Thread:    thread,
		})
	}
}
