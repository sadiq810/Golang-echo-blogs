package api

import (
	"echo_blogs/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type NewBlogController struct {
	DB *gorm.DB
}

func (controller NewBlogController) Index(c echo.Context) error {

	blogRepository := repositories.NewBlogRepository{
		DB: controller.DB,
	}

	blogs := blogRepository.PaginatedList(c)

	return c.JSON(http.StatusOK, blogs)
}

func (controller NewBlogController) Single(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	blogRepository := repositories.NewBlogRepository{
		DB: controller.DB,
	}

	blog := blogRepository.ById(id)

	if blog.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"status": 0, "message": "Blog not found", "data": nil})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": 1, "message": "Blog found", "data": blog})
}
