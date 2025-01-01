package bots

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
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
	//go:embed list.gohtml
	listPageContent  string
	listPageTemplate = layout.MustParse(listPageContent)

	//go:embed assistant.gohtml
	assistantPageContent  string
	assistantPageTemplate = layout.MustParse(assistantPageContent)

	//go:embed create.gohtml
	createPageContent  string
	createPageTemplate = layout.MustParse(createPageContent)

	//go:embed edit.gohtml
	editPageContent  string
	editPageTemplate = layout.MustParse(editPageContent)

	//go:embed chat.gohtml
	chatPageContent  string
	chatPageTemplate = layout.MustParse(chatPageContent)
)

type Resource struct {
	Model *api.Queries
	AI    *openai.Client
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.listPage)
	r.Get("/create", rs.createPage)
	r.Post("/create", rs.postCreate)
	r.Route("/{botID}", func(r chi.Router) {
		r.Use(rs.provideBot)
		r.Get("/", rs.assistantPage)
		r.Get("/chat", rs.chatRedirectPage)
		r.Get("/edit", rs.editPage)
		r.Post("/edit", rs.postEdit)
		r.Get("/chat/{threadID}", rs.chatPage)
		r.Post("/chat/{threadID}/messages", rs.postMessage)
		r.Get("/chat/{threadID}/runstatus/{runID}", rs.getRunStatus)
	})
	return r
}

func (rs Resource) listPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	botRows, err := rs.Model.AllBots(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, listPageTemplate, struct {
		Layout  layout.Data
		BotRows []api.AllBotsRow
	}{
		Layout:  l,
		BotRows: botRows,
	})
}

func (rs Resource) createPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	render.Execute(w, createPageTemplate, struct {
		Layout layout.Data
	}{
		Layout: l,
	})
}

func (rs Resource) postCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	name := r.FormValue("name")
	if name == "" {
		http.Redirect(w, r, "/bots/create", http.StatusSeeOther)
		return
	}
	instructions := r.FormValue("instructions")

	// models, err := client.ListModels(ctx)
	// if err != nil {
	// 	render.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// for _, m := range models.Models {
	// 	fmt.Printf("%v\n", m)
	// }

	model := "gpt-3.5-turbo"

	asst, err := rs.AI.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        model,
		Name:         &name,
		Instructions: &instructions,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bot, err := rs.Model.CreateBot(ctx, api.CreateBotParams{
		OwnerID:     user.ID,
		AssistantID: asst.ID,
		Name:        name,
		Description: instructions, // TODO: replace this with text from the bot introducing itself
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/bots/%d", bot.ID), http.StatusSeeOther)
}

func (rs Resource) editPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	render.Execute(w, editPageTemplate, struct {
		Layout layout.Data
		Bot    api.Bot
	}{
		Layout: l,
		Bot:    bot,
	})
}

func (rs Resource) postEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)
	user := auth.FromContext(ctx)

	name := r.FormValue("name")
	if name == "" {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}
	instructions := r.FormValue("instructions")
	if instructions == "" {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}

	bot, err := rs.Model.UpdateBotDescription(ctx, api.UpdateBotDescriptionParams{
		OwnerID:     user.ID,
		ID:          bot.ID,
		Name:        name,
		Description: instructions,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = rs.AI.ModifyAssistant(ctx, bot.AssistantID, openai.AssistantRequest{
		Name:         &name,
		Instructions: &instructions,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/bots/%d", bot.ID), http.StatusSeeOther)
}

func (rs Resource) assistantPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	render.Execute(w, assistantPageTemplate, struct {
		Layout  layout.Data
		Bot     api.Bot
		IsOwner bool
	}{
		Layout:  l,
		Bot:     bot,
		IsOwner: bot.OwnerID == l.User.ID,
	})
}

func (rs Resource) chatRedirectPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	bot := botFromContext(ctx)

	threads, err := rs.Model.AssistantThreads(ctx, api.AssistantThreadsParams{
		UserID:      user.ID,
		AssistantID: bot.AssistantID,
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var threadID string
	if len(threads) > 0 {
		threadID = threads[0].ThreadID
	} else {
		thread, err := rs.AI.CreateThread(ctx, openai.ThreadRequest{})
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = rs.Model.CreateThread(ctx, api.CreateThreadParams{
			AssistantID: bot.AssistantID,
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

func (rs Resource) chatPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	userThread, err := rs.Model.UserThreadByID(ctx, api.UserThreadByIDParams{
		UserID:   l.User.ID,
		ThreadID: chi.URLParam(r, "threadID"),
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	thread, err := rs.AI.RetrieveThread(ctx, userThread.ThreadID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messagesList, err := rs.AI.ListMessage(ctx, thread.ID, nil, strptr("desc"), nil, nil, nil)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, chatPageTemplate, struct {
		Layout   layout.Data
		Bot      api.Bot
		Thread   openai.Thread
		Messages []openai.Message
		HasMore  bool
	}{
		Layout:   l,
		Bot:      bot,
		Thread:   thread,
		Messages: messagesList.Messages,
		HasMore:  messagesList.HasMore,
	})
}

func (rs Resource) postMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)

	content := strings.TrimSpace(r.FormValue("message"))
	if content == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadID")

	_, err := rs.AI.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	run, err := rs.AI.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: bot.AssistantID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// trigger update to show user's message in the chat
	w.Header().Add("HX-Trigger", "messagesUpdated")

	render.ExecuteNamed(w, chatPageTemplate, "thinking", struct {
		Bot api.Bot
		Run openai.Run
	}{
		Bot: bot,
		Run: run,
	})
}

func (rs Resource) getRunStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)

	run, err := rs.AI.RetrieveRun(ctx, chi.URLParam(r, "threadID"), chi.URLParam(r, "runID"))
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch run.Status {
	case openai.RunStatusQueued, openai.RunStatusInProgress:
		render.ExecuteNamed(w, chatPageTemplate, "thinking", struct {
			Bot api.Bot
			Run openai.Run
		}{
			Bot: bot,
			Run: run,
		})
	default:
		// the run may or may not have been successful, but at this point, we want to
		// trigger an event to update the chat messages
		w.Header().Add("HX-Trigger", "messagesUpdated")

		thread, err := rs.AI.RetrieveThread(ctx, run.ThreadID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.ExecuteNamed(w, chatPageTemplate, "input", struct {
			Bot    api.Bot
			Thread openai.Thread
		}{
			Bot:    bot,
			Thread: thread,
		})
	}
}
