package utility

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type NewTemplateUtility struct {
	C echo.Context
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, d interface{}, c echo.Context) error {
	data := make(map[string]interface{})

	data["data"] = d
	data["User"] = c.Get("user")

	data = checkAndSetFlashMessage(data, c)

	return t.templates.ExecuteTemplate(w, name, data)
}

func checkAndSetFlashMessage(data map[string]interface{}, c echo.Context) map[string]interface{} {
	errors, _ := session.Get("errors", c)
	success, _ := session.Get("success", c)
	successMessage := success.Flashes()
	errorMessage := errors.Flashes()

	_ = errors.Save(c.Request(), c.Response())
	_ = success.Save(c.Request(), c.Response())

	data["successMessage"] = ""
	data["errorMessage"] = ""

	if len(errorMessage) > 0 {
		data["errorMessage"] = errorMessage[0]
	}

	if len(successMessage) > 0 {
		data["successMessage"] = successMessage[0]
	}

	return data
}

func (ntu NewTemplateUtility) SetContext(c echo.Context) NewTemplateUtility {
	ntu.C = c
	return ntu
}

func (ntu NewTemplateUtility) SetTemplate(name string) {
	temp := []string{"views/admin/_layouts/header.html",
		"views/admin/_layouts/head.html",
		"views/admin/_layouts/scripts.html",
		"views/admin/_layouts/navigation.html",
		"views/admin/_layouts/footer.html",
		"views/admin/_layouts/flash_messages.html",
		name,
	}

	t := &Template{
		templates: template.Must(template.ParseFiles(temp...)),
	}

	ntu.C.Echo().Renderer = t
}
