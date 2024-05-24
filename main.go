package main

import (
	"AuthN-AuthZ/handlers"
	"AuthN-AuthZ/repo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handlers struct {
	RegisterUser *handlers.RegisterUserHandler
}

var Handler *Handlers

func main() {
	router := gin.Default()
	router.POST("/api/register", Handler.RegisterUser.RegisterUser)
	router.POST("api/login")
}

func init() {
	db, err := sqlx.Connect("postgres", "user=user123 dbname=AuthN-AuthZ password=1234")
	if err != nil {
		panic(err)
	}
	registerRepo := repo.NewRegisterUserRepo(db)
	registerUser := handlers.NewRegisterUserHandler(registerRepo)
	Handler.RegisterUser = registerUser
}
