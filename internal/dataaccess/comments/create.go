package comments

import (
	"time"

	"github.com/hj235/cvwo/internal/dataaccess/utils"
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func Create(author string, comment *models.Comment) error {
	// Value verification
	if !utils.IsValidUsername(author) {
		return errors.New("comment author is missing or has been deleted")
	}
	if !utils.IsValidComment(comment.Body) {
		return errors.New("comment body is missing")
	}

	comment.Created = time.Now().Format(time.DateTime)

	db := database.GetDB()

	// Add to database
	query := "INSERT INTO comments (author, thread_id, body, time_created) VALUES(?, ?, ?, ?)"
	if _, err := db.Exec(query, author, comment.ThreadId, comment.Body, comment.Created); err != nil {
		return errors.Wrap(err, "error adding comment")
	}

	return nil
}
