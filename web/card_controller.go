package web

import (
	"context"
	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/rubencaro/cardo_server/db"
)

// CardAddHandler handles the creation of new cards
func CardAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
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
