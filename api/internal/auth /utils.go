package auth

import (
	"errors"

	"github.com/BarakaCaleb/orderflow/internal/db"
	"github.com/BarakaCaleb/orderflow/internal/models"
)

// FindUserByEmail finds a user in the database by their email.
func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
