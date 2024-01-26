package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAuthorController interface {
	CreateAuthor(e echo.Context) error
	GetAuthor(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
}

// to access the methods of service and repo
type AuthorController struct {
	authorSvc domain.IAuthorService
}

func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}
func (as *AuthorController) CreateAuthor(e echo.Context) error {
	fmt.Println("CreateAuthor")
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.AuthorDetail{
		Name: reqAuthor.Name,
	}
	if err := as.authorSvc.CreateAuthor(author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail was created successfully")
}
func (as *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.QueryParam("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	author, err := as.authorSvc.GetAuthors(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, author)
}
func (as *AuthorController) UpdateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.AuthorDetail{
		Name: reqAuthor.Name,
	}
	if err := as.authorSvc.UpdateAuthor(author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "AuthorDetail was updated successfully")
}
func (as *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	if err := as.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "AuthorDetail was deleted successfully")
}
