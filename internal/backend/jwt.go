package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/task-manager-api/internal/types"
)

// ValidateJwt validates based on HMAC signing and secret found in env variables
func ValidateJwt(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}

// CreateJwt makes JWT with HMAC-256 encoding
func CreateJwt(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"name":   user.FirstName + " " + user.LastName,
		"nbf":    time.Now(),
		"userId": user.ID,
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}