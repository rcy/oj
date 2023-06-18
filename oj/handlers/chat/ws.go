package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/handlers/render"
	"oj/models/messages"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type RoomSubscriptions map[string]Subscriptions

func (rs RoomSubscriptions) add(roomID string, conn *websocket.Conn) {
	if rs[roomID] == nil {
		rs[roomID] = Subscriptions{}
	}
	rs[roomID].add(conn)
	log.Printf("subscribed %s -- %s -- %d", conn.LocalAddr().String(), roomID, len(rs[roomID]))
}

func (rs RoomSubscriptions) remove(roomID string, conn *websocket.Conn) {
	if rs[roomID] != nil {
		rs[roomID].remove(conn)
	}
	//log.Printf("unsubscribed %s -- %s -- %d", conn.LocalAddr().String(), roomID, len(rs[roomID]))
}

var subs = RoomSubscriptions{}

var messageTemplate = template.Must(template.ParseFiles("handlers/chat/chat_partials.html"))

func ChatServer(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	roomID := chi.URLParam(r, "roomID")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgrade:", err)
		return
	}

	subs.add(roomID, conn)

	defer func() {
		subs.remove(roomID, conn)

		err := conn.Close()
		if err != nil {
			log.Printf("error closing connection: %s", err)
		}
	}()

eventLoop:
	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			var closeError *websocket.CloseError
			if errors.As(err, &closeError) {
				switch closeError.Code {
				case websocket.CloseGoingAway:
				case websocket.CloseNoStatusReceived:
				default:
					log.Printf("*** Unhandled CloseError %v", closeError)
				}
				break eventLoop
			}
			log.Printf("*** Unknown error %v", err)

			break eventLoop
		}

		log.Printf("Received: %s", messageBytes)

		var messageData struct {
			Body string `json:"chat_message"`
		}

		json.Unmarshal(messageBytes, &messageData)

		message, err := messages.Create(roomID, messageData.Body, user.Username)
		if err != nil {
			log.Printf("error creating message: %s", err)
			continue eventLoop
		}

		log.Printf("%s, %s, templates: %s",
			message.Sender,
			user.Username,
			messageTemplate.DefinedTemplates())

		// XXX move these to build time:
		myMsg, err := render.ExecuteNamedToBytes(messageTemplate, "chat_message_mine", message)
		if err != nil {
			log.Printf("error rendering message: %s", err)
			continue eventLoop
		}
		myMsg = []byte(fmt.Sprintf(`<div id="chat_room" hx-swap-oob="beforeend">%s</div>`, myMsg))

		theirMsg, err := render.ExecuteNamedToBytes(messageTemplate, "chat_message_other", message)
		if err != nil {
			log.Printf("error rendering message: %s", err)
			continue eventLoop
		}
		theirMsg = []byte(fmt.Sprintf(`<div id="chat_room" hx-swap-oob="beforeend">%s</div>`, theirMsg))

		// send message to connections
		for c := range subs[roomID] {
			var msg []byte
			if conn == c {
				msg = myMsg
			} else {
				msg = theirMsg
			}
			err = c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("error sending %s to %v", msg, conn)
				// nothing we can really do without a queuing system with retries, etc
			}
		}
	}
}
