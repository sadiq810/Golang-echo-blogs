package admin

import (
	"echo_blogs/repositories"
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type NewCategoryController struct {
	DB *gorm.DB
}

func (dc NewCategoryController) Index(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/category/index.html")

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func (nc NewCategoryController) List(c echo.Context) error {
	newCategoryRepository := repositories.NewCategoryRepository{
		DB: nc.DB,
	}

	data := newCategoryRepository.Datatable(c)

	return c.JSON(http.StatusOK, data)
}

func (dc NewCategoryController) Create(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/category/create.html")

	return c.Render(http.StatusOK, "create.html", map[string]interface{}{})
}

func (dc NewCategoryController) Save(c echo.Context) error {
	status := 0

	title := c.FormValue("title")

	if c.FormValue("status") != "" {
		status = 1
	}

	newCategoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	_ = newCategoryRepository.Save(title, status)

	newFlashMessage.FlashSuccess("Record created successfully")
	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.categories"))
}

func (dc NewCategoryController) Edit(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/category/edit.html")
	newCategoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	cat := newCategoryRepository.ById(id)

	return c.Render(http.StatusOK, "edit.html", cat)
}

func (dc NewCategoryController) Update(c echo.Context) error {
	status := 0

	title := c.FormValue("title")
	id, _ := strconv.Atoi(c.Param("id"))

	if c.FormValue("status") == "on" {
		status = 1
	}

	newCategoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	_ = newCategoryRepository.Update(id, title, status)

	newFlashMessage.FlashSuccess("Record updated successfully")
	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.categories"))
}

func (dc NewCategoryController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	newCategoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	newCategoryRepository.Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "Record deleted successfully."})
}
