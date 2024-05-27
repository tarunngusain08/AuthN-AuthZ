package repo

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRepo struct {
	db *sqlx.DB
}

func NewRegisterUserRepo(db *sqlx.DB) *RegisterUserRepo {
	return &RegisterUserRepo{
		db: db,
	}
}

func (r *RegisterUserRepo) Register(userDetails *contracts.Register) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user := models.User{
		Username: userDetails.Username,
		Email:    userDetails.Email,
		Password: string(hashedPassword),
	}

	// Insert user into database
	query := `INSERT INTO users (username, email, password) VALUES (:username, :email, :password)`
	_, err = r.db.NamedExec(query, &user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	return nil
}
