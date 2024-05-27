package handlers

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	repo *repo.LoginRepo
}

func NewLoginHandler(repo *repo.LoginRepo) *LoginHandler {
	return &LoginHandler{repo: repo}
}

func (l *LoginHandler) Login(c *gin.Context) {
	body := c.Request.Body
	userDetails := new(contracts.Login)
	err := json.NewDecoder(body).Decode(&userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = l.repo.Login(userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Logged in successfully.")
}
