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
	Signup = "users.signup.Signup"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	// REPLACE WITH DATABASE PING CHECK
	// db, err := database.GetDB()

	// if err != nil {
	// 	return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, Signup))
	// }

	var response = api.Response{}
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrParseForm, Signup)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}
	defer r.Body.Close()

	userSensitive, err := usersPkg.Signup(&user)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrCreateFailure, Subject, Signup)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	data, err := json.Marshal(userSensitive)
	if err != nil {
		errorMessage := fmt.Sprintf(msgsPkg.ErrEncodeView, Signup)
		wrappedError := utils.PrepareErrorResponse(&response, err, errorMessage, 1)
		fmt.Println(wrappedError)
		w.WriteHeader(400)
		return &response, wrappedError
	}

	response.Payload.Data = data
	response.Messages = append(response.Messages, fmt.Sprintf(msgsPkg.SuccessfulCreateMessage, Subject))

	return &response, nil
}
