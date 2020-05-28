package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fsunset/twittApp/middlewares"
	"github.com/fsunset/twittApp/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers sets the default-port & listen/serve it
func Handlers() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}

	router := mux.NewRouter()

	// Set "register" path for routers.Register endpoint
	router.HandleFunc(
		"/register", middlewares.CheckingDB(routers.Register),
	).Methods("POST")
	// Set "login" path for routers.Login endpoint
	router.HandleFunc(
		"/login", middlewares.CheckingDB(routers.Login),
	).Methods("POST")
	// Set "profile" path for routers.Profile endpoint
	router.HandleFunc(
		"/profile", middlewares.CheckingDB(
			middlewares.CheckingJWT(routers.Profile),
		),
	).Methods("GET")

	// Listen/Serve PORT
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
