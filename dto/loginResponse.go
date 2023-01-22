package dto

type LoginResponse struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Slug  string  `json:"slug"`
	Token *string `json:"token"`
	Role  string  `json:"role"`
}
