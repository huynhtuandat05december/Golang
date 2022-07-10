package controllers

import (
	"golang_api/dto"
	"golang_api/helpers"
	"golang_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) Profile(ctx *gin.Context) {
	userID := ctx.MustGet("userID")
	convertedUserID, err := strconv.ParseUint(userID.(string), 10, 64)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	profileUser := controller.userService.GetProfileUser(convertedUserID)
	res := helpers.BuildResponse(true, "OK", profileUser)
	ctx.JSON(http.StatusOK, res)

}

func (controller *userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userID := ctx.MustGet("userID")
	convertedUserID, err := strconv.ParseUint(userID.(string), 10, 64)
	if err != nil {
		res := helpers.BuildErrorResponse("Error Server", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userUpdateDTO.ID = convertedUserID
	newProfile := controller.userService.UpdateProfileUser(userUpdateDTO)
	res := helpers.BuildResponse(true, "OK", newProfile)
	ctx.JSON(http.StatusOK, res)

}
