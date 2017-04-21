package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// parseJSONBody receives the Request and returns a map
// with the contents of the Request body. Else panics.
func parseJSONBody(r *http.Request) map[string]interface{} {
	var t map[string]interface{}
	err2 := json.NewDecoder(r.Body).Decode(&t)
	if err2 != nil {
		panic(err2)
	}
	log.Printf("Data read from JSON body: %v\n", t)
	return t
}
