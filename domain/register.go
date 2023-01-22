package domain

import "time"

type Register struct {
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Company     int    `db:"company_id"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Division    string `db:"division"`
	Function    string `db:"function"`
	Designation string `db:"designation"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func NewRegister() Register {
	return Register{
		Name:        "",
		Slug:        "",
		Company:     0,
		Email:       "",
		Password:    "",
		Division:    "vendor",
		Function:    "vendor",
		Designation: "vendor",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}
func UpdateRegister() Register {
	return Register{
		Name:        "",
		Slug:        "",
		Company:     0,
		Email:       "",
		Password:    "",
		Division:    "vendor",
		Function:    "vendor",
		Designation: "vendor",
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
}
