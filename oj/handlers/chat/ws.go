package chat

import (
	"encoding/json"
	"errors"
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

var messageTemplate = template.Must(template.New("").Parse(`<div id="chat_room" hx-swap-oob="beforeend"><div>{{.Sender}}: {{.Body}}</div></div>`))

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
		messageType, messageBytes, err := conn.ReadMessage()
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

		message, err := messages.Create(messageData.Body, user.Username)
		if err != nil {
			log.Printf("error creating message: %s", err)
			continue eventLoop
		}

		outMsg, err := render.ExecuteToBytes(messageTemplate, message)
		if err != nil {
			log.Printf("error creating message: %s", err)
			continue eventLoop
		}

		for conn := range subs[roomID] {
			err = conn.WriteMessage(messageType, outMsg)
			if err != nil {
				log.Println("Error during message writing:", err)
			}
		}
	}
}
