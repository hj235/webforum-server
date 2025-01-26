package users

import (
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) (*models.UserSensitive, error) {
	// Value verification
	if len(username) <= 0 {
		return nil, errors.New("username cannot be empty")
	}
	if len(password) <= 0 {
		return nil, errors.New("password cannot be empty")
	}

	// Retrieve user from database
	retrievedUser, err := getUser(username)
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve the specified user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(password)); err != nil {
		return nil, errors.Wrap(err, "could not verify password")
	}

	// TODO: Implement jwt

	userSensitive := models.UserSensitive{
		Name: retrievedUser.Name,
		Date: retrievedUser.Date,
	}

	return &userSensitive, nil
}
