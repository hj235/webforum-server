package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hj235/cvwo/internal/api"
	usersPkg "github.com/hj235/cvwo/internal/dataaccess/users"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
	"github.com/hj235/cvwo/internal/models"
)

const (
	Delete = "users.delete.Delete"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Delete)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	err = usersPkg.Delete(&user)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrDeleteFailure, Subject, Delete)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	// data, err := json.Marshal(userSensitive)
	// if err != nil {
	// 	errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, Delete)
	// 	return &response, utils.PrepareErrorResponse(&response, err, errorMessage, 1)
	// }

	// response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulDeleteMessage, Subject))

	return &response, nil
}
