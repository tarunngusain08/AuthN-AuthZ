package main

import (
	"AuthN-AuthZ/handlers"
	"AuthN-AuthZ/repo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Handlers struct {
	RegisterUser *handlers.RegisterUserHandler
	LoginUser    *handlers.LoginHandler
	ResetPassword *handlers.ResetPasswordHandler
}

var Handler *Handlers

func main() {
	router := gin.Default()
	router.POST("/api/register", Handler.RegisterUser.RegisterUser)
	router.POST("/api/login", Handler.LoginUser.Login)
	router.POST("/api/reset-password", Handler.ResetPassword.ResetPassword)
	router.Run(":8080")
}

func init() {
	db, err := sqlx.Connect("postgres", "dbname=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	Handler = new(Handlers)
	signingKey := []byte("secret")
	registerRepo := repo.NewRegisterUserRepo(db)
	loginRepo := repo.NewLoginRepo(db, signingKey)

	registerUser := handlers.NewRegisterUserHandler(registerRepo)
	loginHandler := handlers.NewLoginHandler(loginRepo)

	Handler.RegisterUser = registerUser
	Handler.LoginUser = loginHandler
}
