package admin

import (
	"echo_blogs/repositories"
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type NewBlogController struct {
	DB *gorm.DB
}

func (dc NewBlogController) Index(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/blog/index.html")

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func (nc NewBlogController) List(c echo.Context) error {
	blogRepository := repositories.NewBlogRepository{
		DB: nc.DB,
	}

	data := blogRepository.Datatable(c)

	return c.JSON(http.StatusOK, data)
}

func (dc NewBlogController) Create(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/blog/create.html")
	categoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	return c.Render(http.StatusOK, "create.html", map[string]interface{}{"categories": categoryRepository.List()})
}

func (dc NewBlogController) Save(c echo.Context) error {
	title := c.FormValue("title")
	detail := c.FormValue("detail")
	file, _ := c.FormFile("image")

	image := ""

	if file != nil {
		newFileUtility := utility.NewFileUtility{}
		image, _ = newFileUtility.Upload(file, "assets/uploads/")
	}

	categoryId, _ := strconv.Atoi(c.FormValue("category_id"))

	blogRepository := repositories.NewBlogRepository{
		DB: dc.DB,
	}

	_ = blogRepository.Save(title, detail, image, categoryId)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	newFlashMessage.FlashSuccess("Record created successfully")

	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.blogs"))
}

func (dc NewBlogController) Edit(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/blog/edit.html")
	blogRepository := repositories.NewBlogRepository{
		DB: dc.DB,
	}

	categoryRepository := repositories.NewCategoryRepository{
		DB: dc.DB,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	blog := blogRepository.ById(id)

	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{"blog": blog, "categories": categoryRepository.List()})
}

func (dc NewBlogController) Update(c echo.Context) error {
	title := c.FormValue("title")
	detail := c.FormValue("detail")
	categoryId, _ := strconv.Atoi(c.FormValue("category_id"))
	id, _ := strconv.Atoi(c.Param("id"))
	file, _ := c.FormFile("image")

	image := ""

	if file != nil {
		newFileUtility := utility.NewFileUtility{}
		image, _ = newFileUtility.Upload(file, "assets/uploads/")
	}

	blogRepository := repositories.NewBlogRepository{
		DB: dc.DB,
	}

	_ = blogRepository.Update(id, title, detail, image, categoryId)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	newFlashMessage.FlashSuccess("Record updated successfully")

	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.blogs"))
}

func (dc NewBlogController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	blogRepository := repositories.NewBlogRepository{
		DB: dc.DB,
	}

	blogRepository.Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "Record deleted successfully."})
}
