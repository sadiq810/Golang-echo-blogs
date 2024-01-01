package admin

import (
	"echo_blogs/repositories"
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type NewUserController struct {
	DB *gorm.DB
}

type newUserForm struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=5"`
}

type editUserForm struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string
}

func (nuc NewUserController) Index(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/user/index.html")

	return c.Render(http.StatusOK, "index.html", []string{})
}

func (nc NewUserController) List(c echo.Context) error {
	userRepository := repositories.NewUserRepository{
		DB: nc.DB,
	}

	data := userRepository.Datatable(c)

	return c.JSON(http.StatusOK, data)
}

func (dc NewUserController) Create(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/user/create.html")

	return c.Render(http.StatusOK, "create.html", map[string]interface{}{})
}

func (dc NewUserController) Save(c echo.Context) error {
	form := new(newUserForm)

	_ = c.Bind(form)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	if err := c.Validate(form); err != nil {
		errors := newFlashMessage.FormatValidationError(err)

		newFlashMessage.FlashError(errors)
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users.create"))
	}

	userRepository := repositories.NewUserRepository{
		DB: dc.DB,
	}

	userExist := userRepository.ByEmail(form.Email)

	if userExist.ID != 0 {
		newFlashMessage.FlashError("User with provided email already exist")
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users.create"))
	}

	_ = userRepository.Save(form.Name, form.Email, form.Password)

	newFlashMessage.FlashSuccess("Record created successfully")

	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users"))
}

func (dc NewUserController) Edit(c echo.Context) error {
	TemplateUtility.SetContext(c).SetTemplate("views/admin/user/edit.html")

	userRepository := repositories.NewUserRepository{
		DB: dc.DB,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user := userRepository.ById(id)

	return c.Render(http.StatusOK, "edit.html", user)
}

func (dc NewUserController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	form := new(editUserForm)

	_ = c.Bind(form)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	if err := c.Validate(form); err != nil {
		errors := newFlashMessage.FormatValidationError(err)

		newFlashMessage.FlashError(errors)

		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users.edit", id))
	}

	userRepository := repositories.NewUserRepository{
		DB: dc.DB,
	}

	userExist := userRepository.ByEmail(form.Email)

	if userExist.ID != 0 && int(userExist.ID) != id {
		newFlashMessage.FlashError("User with provided email already exist")
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users.edit", id))
	}

	_ = userRepository.Update(id, form.Name, form.Email, form.Password)

	newFlashMessage.FlashSuccess("User updated successfully")
	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.users"))
}

func (dc NewUserController) Delete(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	userRepository := repositories.NewUserRepository{
		DB: dc.DB,
	}

	userRepository.Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "Record deleted successfully."})
}
