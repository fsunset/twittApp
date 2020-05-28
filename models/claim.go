package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim handles JWT structure
type Claim struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
