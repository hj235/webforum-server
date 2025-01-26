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
)

const (
	Subject     = "thread"
	ListThreads = "threads.threads.HandleList"
	GetThread   = "threads.threads.HandleGet"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}

	threads, err := threadsPkg.GetThreads()
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrRetrieveData, ListThreads)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(threads)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, ListThreads)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulListMessage, Subject))

	return &response, nil
}

func HandleGet(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseURLParams, GetThread)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	threads, err := threadsPkg.GetThread(id)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrRetrieveData, GetThread)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(threads)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, GetThread)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulRetrieveMessage, Subject))

	return &response, nil
}
