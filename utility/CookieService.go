package utility

import (
	"echo_blogs/models"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type NewCookieService struct {
	C echo.Context
}

func (ncs NewCookieService) CreateSession(user models.User) {
	sess, _ := session.Get("session", ncs.C)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 4, //7 days
		HttpOnly: true,
	}
	sess.Values["user_id"] = user.ID
	_ = sess.Save(ncs.C.Request(), ncs.C.Response())
}

func (ncs NewCookieService) ExpireUserSession() {
	sess, _ := session.Get("session", ncs.C)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	sess.Values["user_id"] = 0
	_ = sess.Save(ncs.C.Request(), ncs.C.Response())
}

func (ncs NewCookieService) CheckSession(db *gorm.DB) (any, error) {
	sess, _ := session.Get("session", ncs.C)
	user_id := sess.Values["user_id"]

	if user_id == nil || user_id == 0 {
		err := errors.New("Invalid session")
		return nil, err
	}

	user := models.User{}

	db.Model(&models.User{}).Find(&user, map[string]interface{}{"ID": user_id})

	if user.ID == 0 {
		err := errors.New("Invalid session")
		return nil, err
	}

	return user, nil
}
