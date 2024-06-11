package handlers

import (
	"AuthN-AuthZ/contracts"
	"AuthN-AuthZ/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordHandler struct {
	repo *repo.ResetPasswordRepo
}

func NewResetPasswordHandler(repo *repo.ResetPasswordRepo) *ResetPasswordHandler {
	return &ResetPasswordHandler{repo: repo}
}

func (l *ResetPasswordHandler) ResetPassword(c *gin.Context) {
	token := new(contracts.ResetPassword)
	go func(token *contracts.ResetPassword) {
		body := c.Request.Body
		userDetails := new(contracts.ResetPassword)
		err := json.NewDecoder(body).Decode(&userDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = l.repo.ResetPassword(userDetails)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}(token)
	c.JSON(http.StatusOK, token)
}
