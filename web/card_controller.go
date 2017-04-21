package web

import (
	"context"
	"encoding/json"
	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"fmt"
	"github.com/rubencaro/cardo_server/db"
)

// CardAddHandler handles the creation of new cards
func CardAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Data read from Vars: %v\n", vars)

	err1 := r.ParseForm()
	if err1 != nil {
		log.Println(err1)
	}
	log.Printf("Data read from regular POST: %v\n", r.Form)

	var t map[string]interface{}
	err2 := json.NewDecoder(r.Body).Decode(&t)
	if err2 != nil {
		panic(err2)
	}
	log.Printf("Data read from JSON body: %v\n", t)

	Messages <- fmt.Sprintf("%v", t["msg"])
	w.Write([]byte("OKey"))
}

// CardListHandler handles the creation of new cards
func CardListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	ctx := context.Background()
	query := "FOR d IN cardo_card_collection LIMIT 10 RETURN d"
	cursor, err := db.Database.Query(ctx, query, nil)
	if err != nil {
		log.Printf("Error running query '%s': %v", query, err)
	}
	defer cursor.Close()
	for {
		var doc interface{}
		meta, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Printf("Error reading from cursor: %v\n\nMetadata: %v", err, meta)
		}
		log.Println(doc)
	}

	w.Write([]byte("OKey"))
}
