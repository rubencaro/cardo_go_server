package web

import (
	"fmt"
	"log"
	"net/http"
)

// Messages is the channel where you should send your strings
// They will be sent to the client connected
var Messages = make(chan string)

// EventsHandler handles the connection and dispatching
// of Card related events to the client
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "cannot stream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	cn, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "cannot stream", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case <-cn.CloseNotify():
			log.Println("done: closed connection")
			return
		case msg := <-Messages:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			f.Flush()
		}
	}
}
