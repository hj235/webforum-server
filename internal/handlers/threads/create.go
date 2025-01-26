package threads

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hj235/cvwo/internal/api"
	threadsPkg "github.com/hj235/cvwo/internal/dataaccess/threads"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
	"github.com/hj235/cvwo/internal/models"
)

const (
	Create = "threads.create.Create"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}
	thread := models.Thread{}

	// Retrieve URL params (TODO: retrieve from jwt instead?)
	username := chi.URLParam(r, "username")

	// Decode thread information from request body
	err := json.NewDecoder(r.Body).Decode(&thread)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Create)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	err = threadsPkg.Create(username, &thread)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrCreateFailure, Subject, Create)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(thread)
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
