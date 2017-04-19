/*
Package db provides an API to talk to te underlying ArangoDB server
*/
package db

import (
	"fmt"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/rubencaro/cardo_server/helpers"
)

// Connect connects with the Arango instance
func Connect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"https://localhost:8529"},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		// Handle error
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("user", "password"),
	})
	if err != nil {
		// Handle error
	}
}
