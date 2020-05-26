package main

import (
	"log"

	"github.com/fsunset/twittApp/database"
	"github.com/fsunset/twittApp/handlers"
)

func main() {
	if !database.CheckConnection() {
		log.Fatal("Error connecting to MongoDB")
		return
	}

	// HINT --> Avoiding to set directly here Handlers(), so we don't load the whole routes-connections everytime app starts
	handlers.Handlers()
}
