package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fsunset/twittApp/database"
	"github.com/fsunset/twittApp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register sets new users into DB
func Register(w http.ResponseWriter, r *http.Request) {

	// Summon User structure
	var usr models.User

	// Decode info from request's body and sets it into "usr" variable
	err := json.NewDecoder(r.Body).Decode(&usr)

	// If existent; return error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Checking User's email/password data
	if len(usr.Password) == 0 || len(usr.Password) < 6 {
		http.Error(w, "User password cannot be null nor less than 6 chars", 400)
		return
	} else if len(usr.Email) == 0 {
		http.Error(w, "User E-mail cannot be null", 400)
		return
	}

	// Check for email on database; this field must me unique
	_, userFound, _ := database.CheckExistentUser(usr.Email)
	if !!userFound {
		http.Error(w, "Email already exists", 400)
		return
	}

	// Inserting into DB
	_, err = insert(usr)
	if err != nil {
		http.Error(w, "Error inserting new User : "+err.Error(), 400)
		return
	}

	// Return success message response
	w.WriteHeader(http.StatusCreated)
}

func insert(usr models.User) (string, error) {

	// Setting a 10 secs limit to already running DB-context (background)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Cancelling context just before exiting this function
	defer cancel()

	// Encrypts password using default bcrypt-cost of 10
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	usr.Password = string(encryptedPassword)

	// Preparing Database & Collection Constants for this file, into which the insert will be done
	db := database.MongoConnection.Database("twittappcluster")
	collection := db.Collection("users")

	// Inserting...
	result, err := collection.InsertOne(ctx, usr)

	if err != nil {
		return "", err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), nil
}
