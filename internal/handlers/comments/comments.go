package comments

import (
	"encoding/json"
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
	Subject      = "comment"
	ListComments = "comments.comments.HandleList"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}

	// Retrieve URL params
	idParam := chi.URLParam(r, "id")
	fmt.Println("TEST: " + idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseURLParams, ListComments)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	// Data access
	comments, err := commentsPkg.GetComments(id)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrRetrieveData, ListComments)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	// Encode data
	data, err := json.Marshal(comments)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, ListComments)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulListMessage, Subject))

	return &response, nil
}
