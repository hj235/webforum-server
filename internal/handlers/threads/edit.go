package threads

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hj235/cvwo/internal/api"
	threadsPkg "github.com/hj235/cvwo/internal/dataaccess/threads"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
	"github.com/hj235/cvwo/internal/models"
)

const (
	Edit = "edit.Edit"
)

func HandleEdit(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}
	thread := models.Thread{}

	err := json.NewDecoder(r.Body).Decode(&thread)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Edit)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseURLParams, Edit)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	updatedThread, err := threadsPkg.Edit(id, &thread)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEditFailure, Subject, Edit)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(updatedThread)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, Edit)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulEditMessage, Subject))
	return &response, nil
}
