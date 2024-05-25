package main

import (
	"AuthN-AuthZ/handlers"
	"AuthN-AuthZ/repo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Handlers struct {
	RegisterUser *handlers.RegisterUserHandler
}

var Handler *Handlers

func main() {
	router := gin.Default()
	router.POST("/api/register", Handler.RegisterUser.RegisterUser)
	router.POST("api/login")
	router.Run(":8080")
}

func init() {
	db, err := sqlx.Connect("postgres", "dbname=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Handler = new(Handlers)
	registerRepo := repo.NewRegisterUserRepo(db)
	registerUser := handlers.NewRegisterUserHandler(registerRepo)
	Handler.RegisterUser = registerUser
}
