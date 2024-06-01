package repo

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type LoginRepo struct {
	db *sqlx.DB
}

func NewLoginRepo(db *sqlx.DB) *LoginRepo {
	return &LoginRepo{db: db}
}

func (l *LoginRepo) Login(userDetails *contracts.Login) error {
	fetchedUserDetails := new(models.User)
	query := `SELECT * FROM Users WHERE Email = $1`

	err := l.db.Get(fetchedUserDetails, query, userDetails.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(fetchedUserDetails.Password), []byte(userDetails.Password))
	if err != nil {
		return err
	}
	return nil
}
