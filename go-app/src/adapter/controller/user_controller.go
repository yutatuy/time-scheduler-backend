package controller

import (
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterByEmail() gin.HandlerFunc {
	return usecase.RegisterUserByEmail()
}

func LoginByEmail() gin.HandlerFunc {
	return usecase.LoginByEmail()
}