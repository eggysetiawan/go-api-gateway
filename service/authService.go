package service

import (
	"github.com/eggysetiawan/go-api-gateway/domain"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/errs"
	"github.com/eggysetiawan/go-api-gateway/logger"
)

type IAuthService interface {
	Login(username string, password string) (*dto.LoginResponse, *errs.Exception)
	Register(request dto.RegisterRequest) *errs.Exception
}

type DefaultAuthService struct {
	repo domain.IAuthRepository
}

func (s DefaultAuthService) Login(username string, password string) (*dto.LoginResponse, *errs.Exception) {

	login, err := s.repo.FindBy(username, password)

	if err != nil {
		return nil, err
	}

	token, errGenerateToken := login.GenerateToken()

	if errGenerateToken != nil {
		return nil, err
	}

	response := login.ToDto(token)

	return &response, nil
}

func (s DefaultAuthService) Register(request dto.RegisterRequest) *errs.Exception {
	r := domain.NewRegister()

	errEmailExists := s.repo.EmailExists(request.Email)

	if errEmailExists != nil {
		return errEmailExists
	}

	ph, passwordError := request.PasswordValidated()

	if passwordError != nil {
		logger.Error(passwordError.Message)
		return passwordError
	}

	r.Name = request.Name
	r.Slug = request.Slug()
	r.Company = request.Company
	r.Email = request.Email
	r.Password = ph

	err := s.repo.Save(r)

	if err != nil {
		logger.Error(err.Message)
		return err
	}

	return nil
}

func NewAuthService(repository domain.IAuthRepository) DefaultAuthService {
	return DefaultAuthService{repository}
}
