package resources

import (
	"echo_blogs/models"
)

type Category map[string]interface{}

type CategoryResource struct {
}

func (CategoryResource) Resource(i models.Category) Category {

	return Category{
		"id":    i.ID,
		"title": i.Title,
		//"status":     i.Status,
		//"created_at": i.CreatedAt.format("2006-02-01")
	}
}

func (cr CategoryResource) Collection(collection []models.Category) []Category {
	var list []Category

	for _, i := range collection {

		cat := cr.Resource(i)

		list = append(list, cat)
	}

	return list
}
