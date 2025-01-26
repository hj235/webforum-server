package comments

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hj235/cvwo/internal/api"
	commentsPkg "github.com/hj235/cvwo/internal/dataaccess/comments"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
)

const (
	Delete = "delete.Delete"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseURLParams, Delete)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	err = commentsPkg.Delete(id)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrDeleteFailure, Subject, Delete)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulDeleteMessage, Subject))

	return &response, nil
}
