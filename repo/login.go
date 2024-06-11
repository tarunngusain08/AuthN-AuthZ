package repo

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginRepo struct {
	db         *sqlx.DB
	signingKey []byte
}

func NewLoginRepo(db *sqlx.DB, signingKey []byte) *LoginRepo {
	return &LoginRepo{db: db, signingKey: signingKey}
}

func (l *LoginRepo) Login(userDetails *contracts.LoginRequest) (*contracts.LoginResponse, error) {
	fetchedUserDetails := new(models.User)
	query := `SELECT * FROM Users WHERE Email = $1`

	err := l.db.Get(fetchedUserDetails, query, userDetails.Email)
	if err != nil {
		return nil, err
	}
	token := new(string)
	tokenGenerationError := new(error)
	go l.generateToken(token, tokenGenerationError, userDetails)
	err = bcrypt.CompareHashAndPassword([]byte(fetchedUserDetails.Password), []byte(userDetails.Password))
	if err != nil {
		return nil, err
	}
	return &contracts.LoginResponse{Token: token}, *tokenGenerationError
}

func (l *LoginRepo) generateToken(jwtToken *string, tokenGenerationError *error, userDetails *contracts.LoginRequest) {
	// Define the token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      userDetails.Email,
		"password":   userDetails.Password,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(l.signingKey)
	jwtToken = &tokenString
	tokenGenerationError = &err
}
