package users

import (
	"log"

	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Delete(user *models.User) error {
	// Value verification
	if len(user.Name) <= 0 {
		return errors.New("username cannot be empty")
	}
	if len(user.Password) <= 0 {
		return errors.New("password cannot be empty")
	}

	db := database.GetDB()

	// Retrieve user from database
	retrievedUser, err := getUser(user.Name)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password)); err != nil {
		return errors.Wrap(err, "could not verify password")
	}

	// TODO: Implement jwt

	// Delete from database
	query := "DELETE FROM users WHERE username=? AND password=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.Password); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
