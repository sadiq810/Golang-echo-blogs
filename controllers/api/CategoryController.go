package api

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

func (controller NewCategoryController) Index(c echo.Context) error {

	categoryRepository := repositories.NewCategoryRepository{
		DB: controller.DB,
	}

	categories := categoryRepository.PaginatedList(c)

	return c.JSON(http.StatusOK, categories)
}

func (controller NewCategoryController) Single(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	categoryRepository := repositories.NewCategoryRepository{
		DB: controller.DB,
	}

	category := categoryRepository.ById(id)

	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": 0, "message": "Category not found", "data": nil})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": 1, "message": "Category found", "data": category})
}
