/*
Package db provides an API to talk to te underlying ArangoDB server
*/
package db

import (
	"context"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"

	"github.com/rubencaro/cardo_server/hlp"
)

// by now the internal fixed database
var database driver.Database

func init() {
	connect()
}

// connect connects with the Arango, then creates a client
// and sets the 'database' variable to a working instance
func connect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{hlp.Conf.DB.URL},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication(
			hlp.Conf.DB.User,
			hlp.Conf.DB.Pass,
		),
	})
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	ctx := context.Background()
	db, err := client.Database(ctx, "cardo_dev")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	database = db
}

// Coll gets the name of the collection and returns it
func Coll(colname string) driver.Collection {
	col, err := database.Collection(nil, colname)
	if err != nil {
		log.Fatalf("Failed to open collection: %v", err)
	}
	return col
}
