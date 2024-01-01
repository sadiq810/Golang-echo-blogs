package middleware

import (
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type NewCheckAuthMiddleware struct {
	DB *gorm.DB
}

func (cam NewCheckAuthMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCookieService := utility.NewCookieService{C: c}
		_, err := newCookieService.CheckSession(cam.DB)

		if err != nil {
			return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.login"))
		}

		//c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")

		return next(c)
	}
}
