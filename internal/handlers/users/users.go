package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hj235/cvwo/internal/api"
	usersPkg "github.com/hj235/cvwo/internal/dataaccess/users"
	msgsPkg "github.com/hj235/cvwo/internal/handlers/messages"
	"github.com/hj235/cvwo/internal/handlers/utils"
)

const (
	Subject   = "user"
	ListUsers = "users.users.HandleList"
)

func HandleListAll(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}

	users, err := usersPkg.GetUsersSensitive()
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrRetrieveData, ListUsers)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(users)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, ListUsers)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulListMessage, Subject))

	return &response, nil
}
