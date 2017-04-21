package web

import (
	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"fmt"
	"github.com/rubencaro/cardo_server/db"
)

// CardAddHandler handles the creation of new cards
func CardAddHandler(w http.ResponseWriter, r *http.Request) {
	t := parseJSONBody(r)

	_, err := db.Coll("cardo_card_collection").CreateDocument(nil, t)
	if err != nil {
		log.Fatalf("Failed to create document: %v\n%v", t, err)
	}

	Messages <- fmt.Sprintf("Created document: %v", t)
	w.Write([]byte("OKey"))
}

// CardListHandler handles the creation of new cards
func CardListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)

	query := "FOR d IN cardo_card_collection LIMIT 10 RETURN d"
	cursor, err := db.Database.Query(nil, query, nil)
	if err != nil {
		log.Printf("Error running query '%s': %v", query, err)
	}
	defer cursor.Close()
	for {
		var doc interface{}
		meta, err := cursor.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Printf("Error reading from cursor: %v\n\nMetadata: %v", err, meta)
		}
		fmt.Fprintf(w, "%v", doc)
	}

	w.Write([]byte(""))
}

// CardUpdateHandler handles the document update operation
func CardUpdateHandler(w http.ResponseWriter, r *http.Request) {
	t := parseJSONBody(r)

	key, ok := t["key"].(string)
	if !ok {
		log.Fatalf("Couldn't read key: %v", t)
	}

	_, err := db.Coll("cardo_card_collection").UpdateDocument(nil, key, t)
	if err != nil {
		log.Fatalf("Failed to update document: %v\n%v", t, err)
	}
}
