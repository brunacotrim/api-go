package controller

import (
	"github.com/brunacotrim/api-go/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Não implementada a inclusão de usuários")
	c.JSON(err.Code, err)
}
