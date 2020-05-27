package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fsunset/twittApp/database"
	"github.com/fsunset/twittApp/models"
	"github.com/fsunset/twittApp/routers/jwt"
)

// Login initializes logIn process for user, with their email/password
func Login(w http.ResponseWriter, r *http.Request) {

	// Add Response-Header. We're using JSON content
	w.Header().Add("Content-Type", "application/json")

	// Set request-data into u variable
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Actually loging in...
	usr, loggedIn := database.LoginUser(u.Email, u.Password)
	if !loggedIn {
		http.Error(w, "Invalid Credentials", 400)
	}

	// Create JWT
	keyJWT, err := jwt.GenerateJWT(usr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Summon correct Response structure
	loginResponse := models.LoginResponse{
		Token: keyJWT,
	}

	// Set Response-Header. We're using JSON content
	w.Header().Set("Content-Type", "application/json")
	// Return OK status
	w.WriteHeader(http.StatusCreated)
	// Encode Response in JSON, in order to return it
	json.NewEncoder(w).Encode(loginResponse)

	// HINT --> creation of cookie from backend side
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "userJWT",
		Value:   keyJWT,
		Expires: expirationTime,
	})
}
