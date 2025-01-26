package users

import (
	"errors"
	"fmt"
	"log"

	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
)

func GetUsersSensitive() ([]models.UserSensitive, error) {
	db := database.GetDB()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []models.UserSensitive

	for rows.Next() {
		user := models.UserSensitive{}
		pwPlaceholder := ""
		if err := rows.Scan(&user.Name, &pwPlaceholder, &user.Date); err != nil {
			log.Println("Error scanning row: ", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func getUser(name string) (*models.User, error) {
	db := database.GetDB()

	query := "SELECT * FROM users WHERE username=?"

	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("no user matches the given username: %v", name)
	}

	var user models.User
	rows.Scan(&user.Name, &user.Password, &user.Date)

	return &user, nil
}

func GetUserSensitive(name string) (*models.UserSensitive, error) {
	db := database.GetDB()

	query := "SELECT * FROM users WHERE username=?"

	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("no user with the indicated username was found")
	}

	var userSensitive models.UserSensitive
	var pwPlaceholder string
	rows.Scan(&userSensitive.Name, &pwPlaceholder, &userSensitive.Date)

	return &userSensitive, nil
}
