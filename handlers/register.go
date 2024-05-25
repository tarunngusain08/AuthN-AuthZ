package handlers

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterUserHandler struct {
	repo *repo.RegisterUserRepo
}

func NewRegisterUserHandler(repo *repo.RegisterUserRepo) *RegisterUserHandler {
	return &RegisterUserHandler{repo: repo}
}

func (r *RegisterUserHandler) RegisterUser(c *gin.Context) {
	body, err := c.Request.GetBody()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Details Entered")
		return
	}

	var userDetails *contracts.RegisterUser
	err = json.NewDecoder(body).Decode(userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Details Entered")
		return
	}

	err = r.repo.Register(userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not register the user!")
		return
	}

	c.JSON(http.StatusCreated, "User Successfully Registered")
}
