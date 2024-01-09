package controller

import (
	"fmt"
	"net/http"

	"github.com/AhmadAl-Zein/aqary_test/database"
	"github.com/AhmadAl-Zein/aqary_test/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GenerateOTP(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user database.CreateUserParams
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateUser(ctx, user)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Phone Number already exist"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *userController) GenerateOTP(ctx *gin.Context) {
	var req database.UpdateOTPParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.GenerateOTP(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Phone Number doesn't exist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP generated successfully"})
}

func (c *userController) VerifyOTP(ctx *gin.Context) {
	var req database.VerifyOTPParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.VerifyOTP(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
