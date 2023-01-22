package domain

type AccessTokenClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Role string `json:"role"`
}
