package controller

import (
	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
