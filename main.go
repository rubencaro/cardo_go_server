/*
The Cardo Server application is the Go version of the original idea
written in Elixir.
*/

package main

import (
	"net/http"

	_ "github.com/rubencaro/cardo_server/db"
)

func main() {
	// msgBroker = NewBroker()
	// http.HandleFunc("/update", messageHandler)
	// http.HandleFunc("/events", timerEventSource)

	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		panic(err)
	}
}
