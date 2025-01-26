package threads

import (
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func GetThreads() ([]models.Thread, error) {
	db := database.GetDB()

	query := "SELECT * FROM threads"
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve threads")
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		thread := models.Thread{}
		if err := rows.Scan(&thread.Id, &thread.Author, &thread.Title, &thread.Body, &thread.Created, &thread.Edited, &thread.Tags); err != nil {
			return nil, errors.Wrap(err, "unable to scan threads")
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func GetThread(id int) (*models.Thread, error) {
	db := database.GetDB()

	query := "SELECT * FROM threads WHERE id=?"
	row := db.QueryRow(query, id)

	var thread models.Thread
	err := row.Scan(&thread.Id, &thread.Author, &thread.Title, &thread.Body, &thread.Created, &thread.Edited, &thread.Tags)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve thread")
	}

	return &thread, nil
}
