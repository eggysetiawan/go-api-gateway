package app

import (
	"encoding/json"
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/service"
	"net/http"
)

type AuthHandler struct {
	service service.IAuthService
}

//func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
//	return
//}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest

	json.NewDecoder(r.Body).Decode(&request)

	err := ah.service.Register(request)

	if err != nil {
		response := config.NewUnexpectedResponse(err.Message)
		response.Code = err.Code

		config.JsonResponse(w, response.Code, response)

		return
	}

	response := config.NewCreatedResponse("User telah berhasil didaftarkan")

	config.JsonResponse(w, response.Code, response)

	return
}
