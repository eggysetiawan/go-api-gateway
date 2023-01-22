package domain

import (
	"database/sql"
	"github.com/eggysetiawan/go-api-gateway/errs"
	"github.com/eggysetiawan/go-api-gateway/logger"
	"github.com/eggysetiawan/go-api-gateway/support"
	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) EmailExists(e string) *errs.Exception {
	var r Register

	findByEmailSql := "SELECT email FROM users WHERE email = ?"

	err := d.client.Get(&r, findByEmailSql, e)
	if err != nil {

		if err != sql.ErrNoRows {
			return errs.NewUnexpectedException("Error while checking unique email " + err.Error())
		}
		return nil

	}
	return errs.NewUnprocessableEntityException("Email sudah ada sebelumnya")
}

func (d AuthRepositoryDb) PasswordMatch(rp string, dbp string) *errs.Exception {
	err := support.CheckPasswordHash(rp, dbp)

	if err != nil {
		return errs.NewUnprocessableEntityException("Password anda salah!")
	}

	return nil
}

func (d AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.Exception) {
	findBySql := `SELECT users.id, users.name, users.slug, users.password, roles.name as roleName FROM users LEFT OUTER JOIN model_has_roles ON users.id = model_has_roles.model_id LEFT OUTER JOIN roles ON model_has_roles.role_id = roles.id WHERE users.email = ?`

	var login Login

	err := d.client.Get(&login, findBySql, username)

	if err != nil {
		return nil, errs.NewUnexpectedException("error while querying users " + err.Error())
	}

	errPasswordMatch := d.PasswordMatch(password, login.Password)

	if errPasswordMatch != nil {
		return nil, errPasswordMatch
	}

	return &login, nil

}

func (d AuthRepositoryDb) Save(r Register) *errs.Exception {
	registerSql := "INSERT INTO users(`name`,`slug`,`company_id`,`email`,`password`,`division`,`function`,`designation`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)"

	_, err := d.client.Exec(registerSql, r.Name, r.Slug, r.Company, r.Email, r.Password, r.Division, r.Function, r.Designation, r.CreatedAt, r.UpdatedAt)

	if err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedException(err.Error())
	}

	return nil
}

func NewAuthRepositoryDb(db *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{db}
}
