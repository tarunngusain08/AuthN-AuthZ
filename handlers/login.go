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
	token := new(contracts.LoginResponse)
	go func(token *contracts.LoginResponse) {
		body := c.Request.Body
		userDetails := new(contracts.LoginRequest)
		err := json.NewDecoder(body).Decode(&userDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		token, err = l.repo.Login(userDetails)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}(token)
	c.JSON(http.StatusOK, token)
}
