package validator

const validationErrorMessage = "validation request failed"

type ValidationError struct {
	HttpStatusCode int                 `json:"-"`
	Message        string              `json:"message"`
	Internal       error               `json:"-"`
	Details        []map[string]string `json:"details"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) Unwrap() error {
	return e.Internal
}

func (e *ValidationError) Code() int {
	return e.HttpStatusCode
}
