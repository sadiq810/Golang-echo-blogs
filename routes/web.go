package routes

import (
	"echo_blogs/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Web(e *echo.Echo, db *gorm.DB) {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		categoryRepository := repositories.NewCategoryRepository{DB: db}
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"name": "Khan G", "categories": categoryRepository.List()})
	})

	e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is home of the frontend for blogs")
	})
}
