package convertor

import (
	"fmt"
	"go-app/src/application/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginByEmailRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=20"`
}

func NewLoginByEmailConvertor(c *gin.Context) (*usecase.LoginByEmailUsecaseInput, error) {
	var req LoginByEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: %s\n", err)
		}
		return nil, err
	}

	return &usecase.LoginByEmailUsecaseInput{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}