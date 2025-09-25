package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Config
	apiBase := "/api/v1/"

	// Register handlers
	RegisterServerHandlers(apiBase)

	port := "80"

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
