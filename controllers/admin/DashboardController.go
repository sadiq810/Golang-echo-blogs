package admin

import (
	"echo_blogs/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type NewDashboardController struct {
	DB *gorm.DB
}

func (dc NewDashboardController) Index(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/dashboard.html")

	data := make(map[string]interface{})

	newBlogRepository := repositories.NewBlogRepository{DB: dc.DB}
	newCategoryRepository := repositories.NewCategoryRepository{DB: dc.DB}
	newUserRepository := repositories.NewUserRepository{DB: dc.DB}

	data["total_blogs"] = newBlogRepository.GetCount(c)
	data["total_blog_views"] = newBlogRepository.GetTotalViews(c)
	data["total_categories"] = newCategoryRepository.GetCount(c)
	data["total_users"] = newUserRepository.GetCount(c)

	return c.Render(http.StatusOK, "dashboard.html", data)
}
