package users

import (
	"fmt"
	"time"

	"github.com/hj235/cvwo/internal/dataaccess/utils"
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user *models.User) (*models.UserSensitive, error) {
	// Value verification
	if len(user.Name) <= 0 {
		return nil, errors.New("Username cannot be empty")
	}
	if len(user.Password) <= 0 {
		return nil, errors.New("Password cannot be empty")
	}

	user.Date = time.Now().Format(time.DateTime)

	// Get database
	db := database.GetDB()

	// Verify that name does not already exist
	if utils.UsernameExists(user.Name) {
		return nil, errors.New("Username already exists")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Add to database
	query := "INSERT INTO users (username, password, date_created) VALUES(?, ?, ?)"
	// stmt, err := db.Prepare(query)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer stmt.Close()

	// if _, err := stmt.Exec(user.Name, hashedPw, user.Date); err != nil {
	// 	log.Println(err)
	// }

	if _, err := db.Exec(query, user.Name, hashedPw, user.Date); err != nil {
		return nil, errors.Wrap(err, "error adding user")
	}

	userSensitive := models.UserSensitive{
		Name: user.Name,
		Date: user.Date,
	}

	return &userSensitive, nil
}
