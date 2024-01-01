package middleware

import (
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type NewGuestMiddleware struct{ DB *gorm.DB }

func (ngm NewGuestMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCookieService := utility.NewCookieService{C: c}
		user, _ := newCookieService.CheckSession(ngm.DB)

		if user != nil {
			return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.dashboard"))
		}

		return next(c)
	}
}
