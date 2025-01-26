package comments

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hj235/cvwo/internal/dataaccess/utils"
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func Edit(commentId int, comment *models.Comment) (*models.Comment, error) {
	var keys []string
	var values []string

	if utils.IsValidComment(comment.Body) {
		keys = append(keys, models.CommentBodyKey)
		values = append(values, comment.Body)
	}

	db := database.GetDB()

	// TODO: Implement jwt and authentication

	// Set edited timestamp
	keys = append(keys, models.CommentEditedKey)
	values = append(values, time.Now().Format(time.DateTime))

	// Prepare query arguments
	values = append(values, strconv.Itoa(commentId))
	args := make([]any, len(values))
	for i, v := range values {
		args[i] = v
	}

	// Edit comment in database
	keyStr := strings.Join(keys, "=?, ") + "=?"
	query := fmt.Sprintf("UPDATE comments SET %v WHERE id=?", keyStr)
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "could not edit comment in database")
	}
	// TODO: FIX THIS CHECK
	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("%d rows were affected", rowsAffected))
	}

	updatedComment, err := getComment(commentId)
	if err != nil {
		return nil, errors.Errorf("error retrieving comment after updating: %v", err)
	}

	return updatedComment, nil
}
