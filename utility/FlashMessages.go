package utility

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"strings"
)

type NewFlashMessageService struct {
	C echo.Context
}

func (nfm NewFlashMessageService) FlashSuccess(message string) {
	sess, _ := session.Get("success", nfm.C)
	sess.AddFlash(message)
	_ = sess.Save(nfm.C.Request(), nfm.C.Response())
}

func (nfm NewFlashMessageService) FlashError(message string) {
	sess, _ := session.Get("errors", nfm.C)
	sess.AddFlash(message)
	_ = sess.Save(nfm.C.Request(), nfm.C.Response())
}

func (nfm NewFlashMessageService) FormatValidationError(err error) string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errors = append(errors, fmt.Sprintf("%s field is required", err.Field()))
		case "min":
			errors = append(errors, fmt.Sprintf("%s length must be greater or equal to %s", err.Field(), err.Param()))
		}
	}

	return strings.Join(errors, ", ")
}
