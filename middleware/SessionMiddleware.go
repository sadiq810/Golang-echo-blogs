package middleware

import (
	"echo_blogs/utility"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type NewSessionMiddleware struct {
	DB *gorm.DB
}

func (cam NewSessionMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCookieService := utility.NewCookieService{C: c}
		user, _ := newCookieService.CheckSession(cam.DB)

		c.Set("user", user)

		return next(c)
	}
}
