package routes

import (
	"echo_blogs/configs"
	"echo_blogs/controllers/api"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Api(e *echo.Echo, db *gorm.DB) {

	configs.SetupCors(e)
	g := e.Group("api")

	categoryController := api.NewCategoryController{DB: db}
	blogController := api.NewBlogController{DB: db}

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blogs APIs!")
	})

	g.GET("/categories", categoryController.Index)
	g.GET("/categories/:id", categoryController.Single)

	g.GET("/blogs", blogController.Index)
	g.GET("/blogs/:id", blogController.Single)
}
