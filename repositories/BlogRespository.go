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

type NewBlogRepository struct {
	DB *gorm.DB
}

func (br NewBlogRepository) List(categoryId int) []models.Blog {
	var blogs []models.Blog

	builder := br.DB.Model(models.Blog{})

	if categoryId != 0 {
		builder.Where("category_id = ?", categoryId)
	}

	_ = builder.Find(&blogs)

	return blogs
}

func (br NewBlogRepository) ById(id int) models.Blog {
	var blog models.Blog

	_ = br.DB.Model(&models.Blog{}).Where("id = ?", id).Preload("Category").Find(&blog)

	return blog
}

func (cr NewBlogRepository) PaginatedList(c echo.Context) map[string]interface{} {

	var blogs []models.Blog

	var meta = map[string]interface{}{}

	builder := cr.DB

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	var total int64

	cr.DB.Model(&models.Blog{}).Count(&total)

	meta["current_page"] = page
	meta["last_page"] = math.Ceil(float64(total) / float64(pageSize))
	meta["from"] = offset + 1
	meta["to"] = (page-1)*pageSize + pageSize
	meta["total"] = total

	builder.Scopes(models.Paginate(c.Request())).Preload("Category").Find(&blogs)

	return map[string]interface{}{"data": blogs, "meta": meta}
}

func (cr NewBlogRepository) Save(title string, detail string, image string, category_id int) any {
	var blog models.Blog
	blog.Title = title
	blog.Detail = detail
	blog.Image = image
	blog.CategoryId = uint(category_id)

	_ = cr.DB.Create(&blog)

	return blog
}

func (cr NewBlogRepository) Update(id int, title string, detail string, image string, category_id int) any {
	var blog = map[string]interface{}{"Title": title, "Detail": detail, "CategoryId": category_id}

	if image != "" {
		blog["Image"] = image
	}

	_ = cr.DB.Model(&models.Blog{}).Where("ID = ?", id).Updates(blog)

	return blog
}

func (cr NewBlogRepository) Delete(id int) {
	_ = cr.DB.Delete(&models.Blog{}, id)
}

func (cr NewBlogRepository) Datatable(c echo.Context) any {

	newDatatableUtility := utility.NewDatatableUtility{
		C: c,
	}

	var blogs []models.Blog

	builder := cr.DB.Model(&models.Blog{}).Preload("Category")

	ndu := newDatatableUtility.Of(builder)

	_ = builder.Find(&blogs)

	ndu.List = blogs

	return ndu.Transform().AddColumn("action", func(r map[string]interface{}) string {
		id := fmt.Sprint(r["id"])

		return "<a href='/admin/blogs/" + id + "/edit' class='btn btn-primary'>Edit</a> | <a href='#' data-id='" + id + "' class='btn btn-danger delete'>Delete</a>"
	}).Make()
}

func (nbr NewBlogRepository) GetCount(c echo.Context) int64 {
	var total int64

	nbr.DB.Model(&models.Blog{}).Count(&total)

	return total
}

func (nbr NewBlogRepository) GetTotalViews(c echo.Context) int64 {
	var total int64

	nbr.DB.Model(&models.Blog{}).Select("sum(views) as total").Scan(&total)

	return total
}
