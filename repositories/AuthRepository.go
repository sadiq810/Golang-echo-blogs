package repositories

import (
	"echo_blogs/models"
	"echo_blogs/utility"
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NewAuthRepository struct {
	DB *gorm.DB
}

func (ar NewAuthRepository) Authenticate(email string, password string, c echo.Context) error {
	user := models.User{}

	result := ar.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return errors.New("invalid credentials")
	}

	cookieService := utility.NewCookieService{C: c}

	cookieService.CreateSession(user)

	return nil
}

func (ar NewAuthRepository) Logout(c echo.Context) {
	cookieService := utility.NewCookieService{C: c}
	cookieService.ExpireUserSession()
}

func (ar NewAuthRepository) Register(name string, email string, password string) *gorm.DB {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := models.User{Name: name, Email: email, Password: string(hashedPassword)}
	result := ar.DB.Create(&user)

	return result
}
