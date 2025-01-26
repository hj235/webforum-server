package utils

import (
	"github.com/hj235/cvwo/internal/api"

	"github.com/pkg/errors"
)

func PrepareErrorResponse(response *api.Response, err error,
	errorMsg string, errorCode int) error {

	wrappedError := errors.Wrap(err, errorMsg)

	response.Messages = []string{wrappedError.Error()}
	response.ErrorCode = errorCode

	return wrappedError
}
