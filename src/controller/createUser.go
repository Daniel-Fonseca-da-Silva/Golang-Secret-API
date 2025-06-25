package controller

import (
	"fmt"

	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/configuration/rest_err"
	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
