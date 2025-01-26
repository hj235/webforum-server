package comments

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hj235/cvwo/internal/api"
	commentsPkg "github.com/hj235/cvwo/internal/dataaccess/comments"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
	"github.com/hj235/cvwo/internal/models"
)

const (
	Create = "comments.create.Create"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}
	comment := models.Comment{}

	// Retrieve URL params (TODO: retrieve from jwt instead?)
	username := chi.URLParam(r, "username")

	// Decode comment information from request body
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Create)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	// Data access
	err = commentsPkg.Create(username, &comment)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrCreateFailure, Subject, Create)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	comment.Author = sql.NullString{
		String: username,
		Valid:  true,
	}

	// Encode Data
	data, err := json.Marshal(comment)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, Create)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulCreateMessage, Subject))

	return &response, nil
}
