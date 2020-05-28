package routers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fsunset/twittApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile returns user's profile
func Profile(w http.ResponseWriter, r *http.Request) {

	// Get ID parameter from path
	usrID := r.URL.Query().Get("id")
	if len(usrID) < 1 {
		http.Error(w, "Missing parameter 'id'", 400)
	}

	profileFound, err := getProfileData(usrID)

	if err != nil {
		http.Error(w, "Error retrieving User Profile : "+err.Error(), 400)
		return
	}

	// Set Header to json format
	w.Header().Set("Content-Type", "application/json")

	// Set success for response
	w.WriteHeader(http.StatusFound)

	// Return profile in json format
	json.NewEncoder(w).Encode(profileFound)
}

func getProfileData(usrID string) (models.User, error) {

	// Setting a 10 secs limit to the already running DB-context (background)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Cancelling context just before exiting this function
	defer cancel()

	// Transform User's ID from string into ObjectID format...
	objID, _ := primitive.ObjectIDFromHex(usrID)
	// And then into bson...
	queryCondition := bson.M{"_id": objID}

	// Set structure for profile
	var profile models.User
	// Querying DB for ID and saving result into "existentUser"
	err := UsersCollection.FindOne(ctx, queryCondition).Decode(&profile)

	// Removing "password" info from returned data
	profile.Password = ""

	if err != nil {
		log.Fatal(err.Error())
		return profile, err
	}

	return profile, nil
}
