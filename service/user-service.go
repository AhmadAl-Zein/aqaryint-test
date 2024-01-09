package service

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/AhmadAl-Zein/aqary_test/database"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx *gin.Context, user database.CreateUserParams) error
	GenerateOTP(ctx *gin.Context, req database.UpdateOTPParams) error
	VerifyOTP(ctx *gin.Context, req database.VerifyOTPParams) error
}

type userService struct {
	users   []database.User
	queries *database.Queries
}

func New(queries *database.Queries) UserService {
	return &userService{
		users:   make([]database.User, 0),
		queries: queries,
	}
}

func (service *userService) CreateUser(ctx *gin.Context, user database.CreateUserParams) error {
	existingUser, err := service.queries.GetUserByPhoneNumber(ctx, user.PhoneNumber)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingUser.ID != 0 {
		return sql.ErrNoRows
	}

	_, err = service.queries.CreateUser(ctx, user)
	return err
}

func (service *userService) GenerateOTP(ctx *gin.Context, req database.UpdateOTPParams) error {
	_, err := service.queries.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	fmt.Println(err)
	if err != nil {
		return err
	}

	otp := generateRandomOTP()
	otpExpirationTime := time.Now().Add(1 * time.Minute)

	return service.queries.UpdateOTP(ctx, req, otp, otpExpirationTime)
}

func (service *userService) VerifyOTP(ctx *gin.Context, req database.VerifyOTPParams) error {
	_, err := service.queries.VerifyOTP(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func generateRandomOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)
	otpString := fmt.Sprintf("%04d", otp)

	return otpString
}
