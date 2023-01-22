package dto

import (
	"github.com/eggysetiawan/go-api-gateway/errs"
	"github.com/eggysetiawan/go-api-gateway/support"
	"strings"
)

type RegisterRequest struct {
	Name                 string `json:"name"`
	Company              int    `json:"companyId"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r RegisterRequest) Slug() string {
	return strings.ReplaceAll(r.Name, " ", "-")
}

func (r RegisterRequest) PasswordValidated() (string, *errs.Exception) {
	if r.Password != r.PasswordConfirmation {
		return "", errs.NewUnprocessableEntityException("Konfirmasi password tidak cocok")
	}

	passwordHash, err := support.HashPassword(r.Password)

	if err != nil {
		return "", errs.NewUnexpectedException("Error while hashing password " + err.Error())
	}

	return passwordHash, nil

}
