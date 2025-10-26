package auth

import (
	"encoding/json"
	"net/http"

	"github.com/BarakaCaleb/orderflow/internal/db"
	"github.com/BarakaCaleb/orderflow/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles new user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" || input.FullName == "" {
		http.Error(w, "Full name, email, and password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := FindUserByEmail(db.DB, input.Email)
	if err != nil {
		http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}

	if !CheckPassword(user.Password, input.Password) {
		http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// ResetPasswordHandler handles password reset
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, err := FindUserByEmail(db.DB, input.Email)
	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	hashedPassword, err := HashPassword(input.NewPassword)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	db.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
}

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext equivalent.
// It returns true if the password matches the hash.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
