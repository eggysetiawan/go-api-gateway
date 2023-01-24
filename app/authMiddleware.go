package app

import (
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/domain"
	"github.com/eggysetiawan/go-api-gateway/errs"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo domain.IAuthMiddleware
}

func (a AuthMiddleware) AuthorizationHandler() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				response := config.NewUnauthenticated("Unauthenticated")

				config.JsonResponse(w, response.Code, response)

				return
			}

			token := getTokenFromHeader(authHeader)

			errsToken := VerifyToken(token)

			if errsToken != nil {

				response := config.NewForbiddenResponse("Forbidden")

				config.JsonResponse(w, response.Code, response)

				return
			}

			next.ServeHTTP(w, r)

		})
	}
}

func VerifyToken(token string) *errs.Exception {
	_, err := jwtTokenFromString(token)

	if err != nil {
		return err
	}
	return nil

}

func jwtTokenFromString(tokenString string) (*jwt.Token, *errs.Exception) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_KEY_SECRET), nil
	})

	if err != nil {
		return nil, errs.NewForbiddenException(err.Error())
	}

	return token, nil
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
