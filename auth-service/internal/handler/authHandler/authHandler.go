package authHandler

import (
	"auth-service/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// @Summary Login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.Login true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {

	defaultUsername := os.Getenv("DEFAULT_USERNAME")
	defaultPassword := os.Getenv("DEFAULT_PASSWORD")

	w.Header().Set("Content-Type", "application/json")

	var username, password string

	var login dto.Login
	json.NewDecoder(r.Body).Decode(&login)
	username = login.Username
	password = login.Password

	if username == defaultUsername && password == defaultPassword {
		tokenString, err := createToken(username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func createToken(username string) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
