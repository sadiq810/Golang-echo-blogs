package repositories

import (
	"echo_blogs/models"
	"echo_blogs/utility"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type NewCategoryRepository struct {
	DB *gorm.DB
}

func (cr NewCategoryRepository) List() []models.Category {

	var categories []models.Category

	_ = cr.DB.Model(&models.Category{}).Where("status = 1").Find(&categories)

	return categories
}

func (cr NewCategoryRepository) ById(id int) models.Category {
	var category models.Category

	_ = cr.DB.Where("id=?", id).Find(&category)

	return category
}

func (cr NewCategoryRepository) PaginatedList(c echo.Context) map[string]interface{} {

	var categories []models.Category

	var meta = map[string]interface{}{}

	builder := cr.DB.Where("status = 1")

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	var total int64
	cr.DB.Model(&models.Category{}).Count(&total)

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	meta["current_page"] = page
	meta["last_page"] = math.Ceil(float64(total) / float64(pageSize))
	meta["from"] = offset + 1
	meta["to"] = (page-1)*pageSize + pageSize
	meta["total"] = total

	_ = builder.Scopes(models.Paginate(c.Request())).Find(&categories)

	return map[string]interface{}{"data": categories, "meta": meta}
}

func (cr NewCategoryRepository) Save(title string, status int) any {
	var category models.Category
	category.Title = title
	category.Status = status

	_ = cr.DB.Create(&category)

	return category
}

func (cr NewCategoryRepository) Update(id int, title string, status int) any {
	var category = map[string]interface{}{"Title": title, "Status": status}

	_ = cr.DB.Model(&models.Category{}).Where("ID = ?", id).Updates(category)

	return category
}

func (cr NewCategoryRepository) Delete(id int) {
	_ = cr.DB.Delete(&models.Category{}, id)
}

func (cr NewCategoryRepository) Datatable(c echo.Context) any {

	newDatatableUtility := utility.NewDatatableUtility{
		C: c,
	}

	var categories []models.Category

	builder := cr.DB.Model(&models.Category{})

	ndu := newDatatableUtility.Of(builder)

	_ = builder.Find(&categories)

	ndu.List = categories

	return ndu.Transform().AddColumn("action", func(r map[string]interface{}) string {
		id := fmt.Sprint(r["id"])

		return "<a href='/admin/categories/" + id + "/edit' class='btn btn-primary'>Edit</a> | <a href='#' data-id='" + id + "' class='btn btn-danger delete'>Delete</a>"
	}).Make()
}

func (cr NewCategoryRepository) GetCount(c echo.Context) int64 {
	var total int64

	cr.DB.Model(&models.Category{}).Count(&total)

	return total
}
