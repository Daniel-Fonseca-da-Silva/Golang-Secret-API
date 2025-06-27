package controller

import (
	"net/http"

	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/configuration/logger"
	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/configuration/validation"
	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/controller/model/request"
	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Starting create user",
		zap.String("journey", "create_user"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "create_user"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "create_user"),
	)

	c.JSON(http.StatusOK, response)
}
