package config

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewDefaultResponse() Response {
	return Response{
		Message: http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Data:    []string{},
	}
}

func CreatedResponse() Response {
	return Response{
		Message: http.StatusText(http.StatusCreated),
		Code:    http.StatusCreated,
		Data:    []string{},
	}
}

func NewCreatedResponse(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusCreated,
		Data:    []string{},
	}
}

func NewNotFoundResponse() Response {
	return Response{
		Message: http.StatusText(http.StatusNotFound),
		Code:    http.StatusNotFound,
		Data:    []string{},
	}
}

func NewUnexpectedResponse(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusInternalServerError,
		Data:    []string{},
	}
}

func NewBadRequestResponse(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusBadRequest,
		Data:    []string{},
	}
}

func NewForbiddenResponse(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusForbidden,
		Data:    []string{},
	}
}

func NewUnauthorizedResponse(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusUnauthorized,
		Data:    []string{},
	}
}
func NewUnauthenticated(message string) Response {
	return Response{
		Message: message,
		Code:    http.StatusUnauthorized,
		Data:    []string{},
	}
}

func JsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		panic(err)
	}
}
