package threads

import (
	"time"

	"github.com/hj235/cvwo/internal/dataaccess/utils"
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func Create(author string, thread *models.Thread) error {
	// Value verification
	if !utils.IsValidUsername(author) {
		return errors.New("thread author is missing or has been deleted")
	}
	if !utils.IsValidTitle(thread.Title) {
		return errors.New("thread title is missing")
	}

	thread.Created = time.Now().Format(time.DateTime)

	db := database.GetDB()

	// Add to database
	query := "INSERT INTO threads (author, title, body, time_created, tags) VALUES(?, ?, ?, ?, ?)"
	if _, err := db.Exec(query, author, thread.Title, thread.Body, thread.Created, thread.Tags); err != nil {
		return errors.Wrap(err, "error adding thread")
	}

	return nil
}
