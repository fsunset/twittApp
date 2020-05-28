package database

import (
	"context"
	"time"

	"github.com/fsunset/twittApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Preparing Database & Collection Constants for this file, into which the insert will be done...

// AppDB is the database for our app
var AppDB = MongoConnection.Database("twittAppDB")

// UsersCollection is the "users" collection within DB
var UsersCollection = AppDB.Collection("users")

// CheckExistentUser searches a user in DB by email
func CheckExistentUser(email string) (models.User, bool, string) {

	// Setting a 15 secs limit to already running DB-context (background)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	// Cancelling context just before exiting this function
	defer cancel()

	// Converting user's email into bson format
	queryCondition := bson.M{"email": email}

	// Set structure for profile
	var existentUser models.User
	// Querying DB for email and saving result into "existentUser"
	err := UsersCollection.FindOne(ctx, queryCondition).Decode(&existentUser)

	// Get ID from existing returned user
	usrID := existentUser.ID.Hex()

	if err != nil {
		return existentUser, false, usrID
	}

	return existentUser, true, usrID
}
