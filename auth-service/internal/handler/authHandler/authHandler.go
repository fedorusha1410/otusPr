package authHandler

import (
	"auth-service/internal/dto"
	"auth-service/internal/model/user"
	"auth-service/internal/repository"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Handler struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

// @Summary Login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.Login true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var login dto.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	dbUser, err := h.userRepo.GetByUsername(login.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := dbUser.ComparePassword(login.Password); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := createToken(dbUser.Id, int(dbUser.Role))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// @Summary Signup
// @Description Register new user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.SignUp true "New user info"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /signup [post]
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var signup dto.SignUp
	if err := json.NewDecoder(r.Body).Decode(&signup); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Проверим, нет ли такого юзера уже
	_, err := h.userRepo.GetByUsername(signup.Username)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	newUser := &user.User{
		Name: signup.Username,
		Role: user.Creator,
	}

	if err := newUser.SetPassword(signup.Password); err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if err := h.userRepo.Save(newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	tokenString, err := createToken(newUser.Id, int(newUser.Role))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}


func createToken(userID int, role int) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}