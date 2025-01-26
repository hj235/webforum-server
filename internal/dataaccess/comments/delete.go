package comments

import (
	"github.com/hj235/cvwo/internal/database"
	"github.com/pkg/errors"
)

func Delete(commentId int) error {
	db := database.GetDB()

	// Retrieve comment from database
	// retrievedComment, err := GetComment(commentId)
	_, err := getComment(commentId)
	if err != nil {
		return errors.Wrap(err, "error deleting comment: ")
	}

	// TODO: Implement jwt

	// Delete from database
	query := "DELETE FROM comments WHERE id=?"
	result, err := db.Exec(query, commentId)
	if err != nil {
		return errors.Wrap(err, "error deleting comment: ")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("could not check rows affected")
	}
	if rowsAffected < 1 {
		return errors.Errorf("error deleting comment, %v rows affected", rowsAffected)
	}

	return nil
}
