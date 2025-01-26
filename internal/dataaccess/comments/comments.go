package comments

import (
	"github.com/hj235/cvwo/internal/database"
	"github.com/hj235/cvwo/internal/models"
	"github.com/pkg/errors"
)

func GetComments(id int) ([]models.Comment, error) {
	db := database.GetDB()

	query := "SELECT * FROM comments WHERE thread_id=? ORDER BY time_created DESC"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve comments")
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(&comment.Id, &comment.Author, &comment.ThreadId, &comment.Body, &comment.Created, &comment.Edited); err != nil {
			return nil, errors.Wrap(err, "unable to scan comments")
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func getComment(id int) (*models.Comment, error) {
	db := database.GetDB()

	query := "SELECT * FROM comments WHERE id=?"
	row := db.QueryRow(query, id)

	var comment models.Comment
	err := row.Scan(&comment.Id, &comment.Author, &comment.ThreadId, &comment.Body, &comment.Created, &comment.Edited)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve comment")
	}

	return &comment, nil
}
