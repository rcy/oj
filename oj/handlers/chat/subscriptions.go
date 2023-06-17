package chat

import (
	"github.com/gorilla/websocket"
)

type Subscriptions map[*websocket.Conn]bool

func (s Subscriptions) add(conn *websocket.Conn) {
	s[conn] = true
	//log.Printf("subscribed %s -- %d", conn.LocalAddr().String(), len(s))
}

func (s Subscriptions) remove(conn *websocket.Conn) {
	delete(s, conn)
	//log.Printf("unsubscribing %s -- %d", conn.LocalAddr().String(), len(s))
}
