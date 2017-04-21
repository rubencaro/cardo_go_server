/*
The Cardo Server application is the Go version of the original idea
written in Elixir.
*/

package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	_ "github.com/rubencaro/cardo_server/db"
	"github.com/rubencaro/cardo_server/web"
)

func main() {
	// define routes
	r := mux.NewRouter()
	r.HandleFunc("/ping", web.PingHandler)
	r.HandleFunc("/events", web.EventsHandler)
	r.HandleFunc("/list", web.CardListHandler)

	s := r.Methods("POST").
		Headers("Content-Type", "application/json").
		Subrouter()
	s.HandleFunc("/add", web.CardAddHandler)
	s.HandleFunc("/update", web.CardAddHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	// start server
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
