package eventsource

import (
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
)

var SSE = sse.NewServer(nil)

type Realtime struct{}

func New() Realtime {
	return Realtime{}
}

func (Realtime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	SSE.ServeHTTP(w, r)
}
