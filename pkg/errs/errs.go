package errs

import "net/http"

type MessageErr interface {
	Error() string
	Code() int
}

type errData struct {
	ErrStatus  string `json:"status"`
	ErrMessage any    `json:"message"`
	ErrCode    int    `json:"code"`
}

func (e *errData) Error() string {
	return e.ErrStatus
}

func (e *errData) Code() int {
	return e.ErrCode
}

func NewNotFoundError(message string) MessageErr {
	return &errData{
		ErrStatus:  "NOT_FOUND",
		ErrMessage: message,
		ErrCode:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) MessageErr {
	return &errData{
		ErrStatus:  "INTERNAL_SERVER_ERROR",
		ErrMessage: message,
		ErrCode:    http.StatusInternalServerError,
	}
}

func NewBadRequest(message string) MessageErr {
	return &errData{
		ErrStatus:  "BAD_REQUEST",
		ErrMessage: message,
		ErrCode:    http.StatusBadRequest,
	}
}

func NewUnprocessableEntity(messages []string) MessageErr {
	return &errData{
		ErrStatus:  "BAD_REQUEST",
		ErrMessage: messages,
		ErrCode:    http.StatusBadRequest,
	}
}
