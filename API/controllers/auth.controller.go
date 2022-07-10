package controllers

import (
	"fmt"
	"golang_api/dto"
	"golang_api/helpers"
	"golang_api/models"
	"golang_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
func (controller *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	existingUser := controller.authService.VerifyAccount(loginDTO.Email, loginDTO.Password)
	user, existing := existingUser.(models.User)
	fmt.Print(user, existing)
	if !existing {
		response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return

	}
	accessToken := controller.jwtService.GenerateAccessToken(strconv.FormatUint(user.ID, 10))
	refreshToken := controller.jwtService.GenerateRefreshToken(strconv.FormatUint(user.ID, 10))
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	response := helpers.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)
}

func (controller *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !controller.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		newUser := controller.authService.CreateUser(registerDTO)
		accessToken := controller.jwtService.GenerateAccessToken(strconv.FormatUint(newUser.ID, 10))
		refreshToken := controller.jwtService.GenerateRefreshToken(strconv.FormatUint(newUser.ID, 10))
		newUser.AccessToken = accessToken
		newUser.RefreshToken = refreshToken
		response := helpers.BuildResponse(true, "OK!", newUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
