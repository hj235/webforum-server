package threads

import (
	"github.com/hj235/cvwo/internal/database"
	"github.com/pkg/errors"
)

func Delete(threadId int) error {
	db := database.GetDB()

	// Retrieve thread from database
	// retrievedThread, err := GetThread(threadId)
	_, err := GetThread(threadId)
	if err != nil {
		return errors.Wrap(err, "error deleting thread: ")
	}

	// TODO: Implement jwt

	// Delete from database
	query := "DELETE FROM threads WHERE id=?"
	result, err := db.Exec(query, threadId)
	if err != nil {
		return errors.Wrap(err, "error deleting thread: ")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("could not check rows affected")
	}
	if rowsAffected < 1 {
		return errors.Errorf("error deleting thread, %v rows affected", rowsAffected)
	}

	return nil
}
