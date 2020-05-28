package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fsunset/twittApp/database"
	"github.com/fsunset/twittApp/models"
)

// ActiveUserEmail "global" var within routers package
var ActiveUserEmail string

// ActiveUserID "global" var within routers package
var ActiveUserID string

// ValidateJWT decodes JWT & checks if it is valid
func ValidateJWT(token string) (models.Claim, bool, string, error) {
	myUniqueSecretSignature := []byte("my-256-bit-secret-for-JWT")
	var claims models.Claim

	// Split JWT string
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Bad JWT format")
	}

	// Triming token string after the split
	token = strings.TrimSpace(splitToken[1])

	// Decode & validate token using secret-word
	decodedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myUniqueSecretSignature, nil
	})

	// If token is correct check for user
	if err == nil {
		_, foundUsr, _ := database.CheckExistentUser(claims.Email)

		// After User has been found using their JWT, set these 2 "global" variables
		if !!foundUsr {
			ActiveUserEmail = claims.Email
			ActiveUserID = claims.ID.Hex()
		}

		return claims, foundUsr, ActiveUserID, nil
	}

	// If token was not valid
	if !decodedToken.Valid {
		return claims, false, "", errors.New("Invalid Token")
	}

	// Returning error
	return claims, false, "", err
}
