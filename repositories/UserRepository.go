package repositories

import (
	"echo_blogs/models"
	"echo_blogs/utility"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type NewUserRepository struct {
	DB *gorm.DB
}

func (cr NewUserRepository) List() []models.User {

	var users []models.User

	_ = cr.DB.Model(&models.User{}).Where("status = 1").Find(&users)

	return users
}

func (cr NewUserRepository) ById(id int) models.User {
	var user models.User

	_ = cr.DB.Where("id=?", id).Find(&user)

	return user
}

func (cr NewUserRepository) ByEmail(email string) models.User {
	var user models.User

	_ = cr.DB.Where("email=?", email).Find(&user)

	return user
}

func (cr NewUserRepository) PaginatedList(c echo.Context) map[string]interface{} {

	var users []models.User

	var meta = map[string]interface{}{}

	builder := cr.DB

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	if pageSize != 0 || page != 0 {
		builder.Scopes(models.Paginate(c.Request()))

		var total int64
		cr.DB.Model(&models.User{}).Count(&total)

		if page <= 0 {
			page = 1
		}

		meta["current_page"] = page
		meta["last_page"] = math.Ceil(float64(total / int64(pageSize)))
		meta["from"] = (page - 1) * pageSize
		meta["to"] = (page-1)*pageSize + pageSize
		meta["total"] = total
	}

	_ = builder.Find(&users)

	return map[string]interface{}{"data": users, "meta": meta}
}

func (cr NewUserRepository) Save(name string, email string, password string) any {
	var user models.User
	user.Name = name
	user.Email = email

	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user.Password = string(hashedBytes)

	_ = cr.DB.Create(&user)

	return user
}

func (cr NewUserRepository) Update(id int, name string, email string, password string) any {
	var user = map[string]interface{}{"Name": name, "Email": email}

	if password != "" {
		hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		user["Password"] = string(hashedBytes)
	}

	_ = cr.DB.Model(&models.User{}).Where("ID = ?", id).Updates(user)

	return user
}

func (cr NewUserRepository) Delete(id int) {
	_ = cr.DB.Delete(&models.User{}, id)
}

func (cr NewUserRepository) Datatable(c echo.Context) any {

	newDatatableUtility := utility.NewDatatableUtility{
		C: c,
	}

	var users []models.User

	builder := cr.DB.Model(&models.User{})

	ndu := newDatatableUtility.Of(builder)

	_ = builder.Find(&users)

	ndu.List = users

	return ndu.Transform().AddColumn("action", func(r map[string]interface{}) string {
		id := fmt.Sprint(r["id"])

		return "<a href='/admin/users/" + id + "/edit' class='btn btn-primary'>Edit</a> | <a href='#' data-id='" + id + "' class='btn btn-danger delete'>Delete</a>"
	}).Make()
}

func (cr NewUserRepository) GetCount(c echo.Context) int64 {
	var total int64

	cr.DB.Model(&models.User{}).Count(&total)

	return total
}
