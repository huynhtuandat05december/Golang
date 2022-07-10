package controllers

import (
	"golang_api/dto"
	"golang_api/helpers"
	"golang_api/models"
	"golang_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	// Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}

}

func (controller *bookController) All(ctx *gin.Context) {
	var books []models.Book = controller.bookService.GetAllBook()
	res := helpers.BuildResponse(true, "OK", books)
	ctx.JSON(http.StatusOK, res)

}

func (controller *bookController) FindByID(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	book := controller.bookService.GetBookByID(bookID)
	if (book == models.Book{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := helpers.BuildResponse(true, "OK", book)
	ctx.JSON(http.StatusOK, res)

}

func (controller *bookController) Insert(ctx *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := ctx.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userID := ctx.MustGet("userID")
	convertedUserID, err := strconv.ParseUint(userID.(string), 10, 64)
	if err == nil {
		bookCreateDTO.UserID = convertedUserID
	}
	result := controller.bookService.InsertBook(bookCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusCreated, response)

}

// func (controller *bookController) Update(ctx *gin.Context) {

// }

func (controller *bookController) Delete(ctx *gin.Context) {
	bookID, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
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
	if !controller.bookService.IsAllowedToEdit(convertedUserID, bookID) {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	controller.bookService.DeleteBook(bookID)
	res := helpers.BuildResponse(true, "OK", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
