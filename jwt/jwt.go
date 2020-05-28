package jwt

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fsunset/twittApp/models"
)

// GenerateJWT creates JWT
func GenerateJWT(usr models.User) (string, error) {

	// Create the secret-string part, needed by JWT to generate token
	myUniqueSecretSignature := []byte("my-256-bit-secret-for-JWT")

	// Create the payload part, needed by JWT to generate token
	payload := jwt.MapClaims{
		"email":    usr.Email,
		"name":     usr.Name,
		"lastName": usr.LastName,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	// Create token base, passing Header-Encryption-method & claims
	usrToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Finish token creation, by adding secret-signature
	signedToken, err := usrToken.SignedString(myUniqueSecretSignature)

	if err != nil {
		log.Fatal(err.Error())
		return signedToken, err
	}

	return signedToken, nil
}
