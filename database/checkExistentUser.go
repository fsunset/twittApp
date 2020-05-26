package database

import (
	"context"
	"time"

	"github.com/fsunset/twittApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckExistentUser searches a user in DB by email
func CheckExistentUser(email string) (models.User, bool, string) {

	// Setting a 15 secs limit to already running DB-context (background)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	// Cancelling context just before exiting this function
	defer cancel()

	// Preparing Database & Collection Constants for this file, into which the insert will be done
	db := MongoConnection.Database("twittappcluster")
	collection := db.Collection("users")

	// Converting user's email into bson format
	queryCondition := bson.M{"email": email}

	// Querying DB for email and saving result into "existentUser"
	var existentUser models.User
	err := collection.FindOne(ctx, queryCondition).Decode(&existentUser)

	usrID := existentUser.ID.Hex()

	if err != nil {
		return existentUser, false, usrID
	}

	return existentUser, true, usrID
}
