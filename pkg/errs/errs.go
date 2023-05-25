package errs

import "net/http"

type MessageErr interface {
	Error() string
	Status() int
}

type errData struct {
	ErrStatus  int    `json:"status"`
	ErrMessage any    `json:"message"`
	ErrError   string `json:"error"`
}

func (e *errData) Error() string {
	return e.ErrError
}

func (e *errData) Status() int {
	return e.ErrStatus
}

func NewNotFoundError(message string) MessageErr {
	return &errData{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
		ErrError:   "NOT_FOUND",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &errData{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewBadRequest(message string) MessageErr {
	return &errData{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
		ErrError:   "BAD_REQUEST",
	}
}

func NewUnprocessableEntity(messages []string) MessageErr {
	return &errData{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: messages,
		ErrError:   "BAD_REQUEST",
	}
}
