package errs

import (
	"net/http"
)

type Exception struct {
	Message string
	Code    int
}

func (e Exception) AsMessage() *Exception {
	return &Exception{
		Message: e.Message,
	}
}

func NewException(message string, code int) *Exception {
	return &Exception{
		Message: message,
		Code:    code,
	}
}

// 404
func NewNotFoundException(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NotFoundException() *Exception {
	return &Exception{
		Message: http.StatusText(http.StatusNotFound),
		Code:    http.StatusNotFound,
	}
}

// 401
func NewUnauthorizedException(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}
func UnauthorizedException() *Exception {
	return &Exception{
		Message: http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
	}
}

// 403
func NewForbiddenException(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func ForbiddenException() *Exception {
	return &Exception{
		Message: http.StatusText(http.StatusForbidden),
		Code:    http.StatusForbidden,
	}
}

// 422
func NewValidationError(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

// 500
func InternalServerErrorException() *Exception {
	return &Exception{
		Message: http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
	}
}

func NewUnexpectedException(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// 422
func NewUnprocessableEntityException(message string) *Exception {
	return &Exception{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
