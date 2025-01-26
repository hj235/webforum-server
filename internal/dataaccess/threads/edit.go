package threads

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func Edit(threadId int, thread *models.Thread) (*models.Thread, error) {
	var keys []string
	var values []string

	if len(thread.Title) > 0 {
		keys = append(keys, models.ThreadTitleKey)
		values = append(values, thread.Title)
	}
	if len(thread.Body) > 0 {
		keys = append(keys, models.ThreadBodyKey)
		values = append(values, thread.Body)
	}
	if len(thread.Tags) > 0 {
		keys = append(keys, models.ThreadTagsKey)
		values = append(values, thread.Tags)
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("no fields were changed")
	}

	db := database.GetDB()

	// TODO: Implement jwt and authentication

	// Set edited timestamp
	keys = append(keys, models.ThreadEditedKey)
	values = append(values, time.Now().Format(time.DateTime))

	// Prepare query arguments
	values = append(values, strconv.Itoa(threadId))
	args := make([]any, len(values))
	for i, v := range values {
		args[i] = v
	}

	// Edit thread in database
	keyStr := strings.Join(keys, "=?, ") + "=?"
	query := fmt.Sprintf("UPDATE threads SET %v WHERE id=?", keyStr)
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "could not edit thread in database")
	}
	// TODO: FIX THIS CHECK
	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("%d rows were affected", rowsAffected))
	}

	updatedThread, err := GetThread(threadId)
	if err != nil {
		return nil, errors.Errorf("error retrieving thread after updating: %v", err)
	}

	return updatedThread, nil
}
