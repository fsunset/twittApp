package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers sets the default-port & listen/serve it
func Handlers() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router := mux.NewRouter()

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
