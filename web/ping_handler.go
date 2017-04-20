/*
Package web holds all code related to web handling,
well separated from the actual business logic of the application
*/
package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// PingHandler handles the ping request
func PingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	w.Write([]byte("OKey"))
}
