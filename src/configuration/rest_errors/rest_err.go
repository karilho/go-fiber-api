package rest_errors

import "net/http"

// Struct para retornar o "padrão" de erro
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

// Struct para retornar as CAUSAS do erro gerado
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Work to satisfact the interface error from Go
// Use this for another librarier that implements another types of errors.
func (e *RestErr) Error() string {
	return e.Message
}

// Work like a constructors
func NewRestErr(neessage, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: neessage,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}