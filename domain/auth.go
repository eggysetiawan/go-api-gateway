package domain

import (
	"database/sql"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/errs"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const TOKEN_EXPIRATION = 3 * time.Hour

const HMAC_KEY_SECRET = "verythoughtsecret"

type Login struct {
	Id       string         `db:"id"`
	Name     string         `db:"name"`
	Slug     string         `db:"slug"`
	Password string         `db:"password"`
	RoleName sql.NullString `db:"roleName"`
	Token    *string
}

func (l Login) ToDto() dto.LoginResponse {
	return dto.LoginResponse{
		Id:    l.Id,
		Name:  l.Name,
		Slug:  l.Slug,
		Token: l.Token,
		Role:  l.Role(),
	}
}

func (l Login) Role() string {
	if !l.RoleName.Valid {
		return ""
	}
	return l.RoleName.String
}

func (l Login) GenerateToken() (*string, *errs.Exception) {
	var claims jwt.MapClaims

	claims = l.Claims()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(HMAC_KEY_SECRET))

	if err != nil {
		return nil, errs.NewUnexpectedException(err.Error())
	}

	return &signedString, nil
}

func (l Login) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":         l.Id,
		"name":       l.Name,
		"slug":       l.Slug,
		"role":       l.RoleName,
		"expiration": TOKEN_EXPIRATION,
	}
}

type IAuthRepository interface {
	PasswordMatch(rp string, dp string) *errs.Exception
	FindBy(username string, password string) (*Login, *errs.Exception)
	Save(r Register) *errs.Exception
	EmailExists(e string) *errs.Exception
}

type IAuthMiddleware interface {
	AuthorizationHandler() func(handler http.Handler) http.Handler
}
