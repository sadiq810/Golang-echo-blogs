package controllers

import (
	"echo_blogs/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type NewCategoryController struct {
	DB *gorm.DB
}

func (nc NewCategoryController) Index(c echo.Context) error {
	categoryRepository := repositories.NewCategoryRepository{DB: nc.DB}

	return c.JSON(http.StatusOK, categoryRepository.List())
}

func (nc NewCategoryController) Single(c echo.Context) error {
	categoryRepository := repositories.NewCategoryRepository{DB: nc.DB}

	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, categoryRepository.ById(id))
}
