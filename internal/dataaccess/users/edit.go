package users

import (
	"fmt"
	"strings"

	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func Edit(username string, user *models.User) (*models.UserSensitive, error) {
	var keys []string
	var values []string

	if len(user.Name) > 0 {
		keys = append(keys, models.UserNameKey)
		values = append(values, user.Name)
	}
	if len(user.Password) > 0 {
		keys = append(keys, models.UserPasswordKey)
		values = append(values, user.Password)
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("no fields were changed")
	}

	db := database.GetDB()

	// TODO: Implement jwt and authentication

	// Prepare arguments
	values = append(values, username)
	args := make([]any, len(values))
	for i, v := range values {
		args[i] = v
	}

	// Edit user in database
	keyStr := strings.Join(keys, "=?, ") + "=?"
	query := fmt.Sprintf("UPDATE users SET %v WHERE username=?", keyStr)
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "could not edit user in database")
	}
	// TODO: FIX THIS CHECK
	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("%d rows were affected", rowsAffected))
	}

	row := db.QueryRow("SELECT username, date_created FROM users WHERE username=?", user.Name)

	var updatedUser models.UserSensitive
	if err := row.Scan(&updatedUser.Name, &updatedUser.Date); err != nil {
		return nil, errors.Wrap(err, "could not retrieve updated user")
	}

	return &updatedUser, nil
}
