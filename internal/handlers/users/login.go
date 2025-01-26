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
	Login = "users.login.Login"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	var response = api.Response{}
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Login)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	userSensitive, err := usersPkg.Login(user.Name, user.Password)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrLoginFailure, Login)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(userSensitive)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, Login)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, msgsPkg.SuccessfulLoginMessage)

	return &response, nil
}
