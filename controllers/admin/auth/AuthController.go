package auth

import (
	"echo_blogs/controllers/admin"
	"echo_blogs/repositories"
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type NewAuthController struct {
	DB *gorm.DB
}

type LoginFormData struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type RegisterFormData struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

func (ac NewAuthController) ShowLogin(c echo.Context) error {
	admin.TemplateUtility.SetContext(c).SetTemplate("views/admin/auth/login.html")

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func (ac NewAuthController) Login(c echo.Context) error {
	form := new(LoginFormData)

	_ = c.Bind(form)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	if err := c.Validate(form); err != nil {
		errors := newFlashMessage.FormatValidationError(err)

		newFlashMessage.FlashError(errors)

		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.login"))
	}

	authRepo := repositories.NewAuthRepository{DB: ac.DB}

	err := authRepo.Authenticate(form.Email, form.Password, c)

	if err != nil {
		newFlashMessage.FlashError(err.Error())
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.login"))
	}

	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.dashboard"))
}

func (ac NewAuthController) Logout(c echo.Context) error {

	authRepo := repositories.NewAuthRepository{
		DB: ac.DB,
	}

	authRepo.Logout(c)

	return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.login"))
}

func (ac NewAuthController) ShowRegister(c echo.Context) error {
	admin.TemplateUtility.SetContext(c).SetTemplate("views/admin/auth/register.html")
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
}

func (ac NewAuthController) Register(c echo.Context) error {
	form := new(RegisterFormData)

	newFlashMessage := utility.NewFlashMessageService{
		C: c,
	}

	_ = c.Bind(form)

	if err := c.Validate(form); err != nil {
		errors := newFlashMessage.FormatValidationError(err)

		newFlashMessage.FlashError(errors)
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.register"))
	}

	userRepo := repositories.NewUserRepository{DB: ac.DB}

	userExist := userRepo.ByEmail(form.Email)

	if userExist.ID != 0 {
		newFlashMessage.FlashError("User with provided email already exist")
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.register"))
	}

	authRepo := repositories.NewAuthRepository{DB: ac.DB}

	result := authRepo.Register(form.Name, form.Email, form.Password)

	if result.Error != nil {
		newFlashMessage.FlashError("Error occurred, " + result.Error.Error())
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.register"))
	} else {
		newFlashMessage.FlashSuccess("Your registration was successful. You can login now.")
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.login"))
	}
}
