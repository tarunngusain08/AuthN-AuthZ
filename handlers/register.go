package handlers

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RegisterUserHandler struct {
	repo *repo.RegisterUserRepo
}

func NewRegisterUserHandler(repo *repo.RegisterUserRepo) *RegisterUserHandler {
	return &RegisterUserHandler{repo: repo}
}

func (r *RegisterUserHandler) RegisterUser(c *gin.Context) {
	body := c.Request.Body
	var userDetails *contracts.RegisterUser
	err := json.NewDecoder(body).Decode(&userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Details Entered")
		return
	}

	err = r.repo.Register(userDetails)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.JSON(http.StatusInternalServerError, "User already Exists! Try logging in or use alternate email")
		} else {
			c.JSON(http.StatusInternalServerError, "Could not register the user!"+err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, "User Successfully Registered")
}
