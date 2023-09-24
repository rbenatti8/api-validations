package binder

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
)

const validationErrorMessage = "validation request failed"

type BindError struct {
	HttpStatusCode int                 `json:"-"`
	Message        string              `json:"message"`
	Details        []map[string]string `json:"details"`
}

func (e *BindError) Error() string {
	return e.Message
}

func (e *BindError) Code() int {
	return e.HttpStatusCode
}

func handleErrors(err error) error {
	if jsonErr := new(json.SyntaxError); errors.As(err, &jsonErr) {
		return &BindError{
			Message: validationErrorMessage,
			Details: []map[string]string{
				{
					"body": "the request body is not in a valid json",
				},
			},
		}
	}

	if jsonErr := new(json.UnmarshalTypeError); errors.As(err, &jsonErr) {
		location := fmt.Sprintf("body.%s", jsonErr.Field)
		message := fmt.Sprintf("expected %s but got %s", jsonErr.Type.String(), jsonErr.Value)

		return &BindError{
			Message: validationErrorMessage,
			Details: []map[string]string{
				{
					location: message,
				},
			},
		}
	}

	return err
}
