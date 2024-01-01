package utility

import (
	_ "echo_blogs/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
)

type NewDatatableUtility struct {
	C               echo.Context
	List            interface{}
	draw            int
	recordsTotal    int64
	recordsFiltered int64
}

func (ndu NewDatatableUtility) Of(builder *gorm.DB) NewDatatableUtility {
	qParams := ndu.C.QueryParams()
	orderColumnIndex := qParams["order[0][column]"][0]
	orderDirection := qParams["order[0][dir]"][0]
	orderColumnName := qParams["columns["+orderColumnIndex+"][name]"][0]
	searchStr := qParams["search[value]"][0]

	var columns []string

	for k, v := range qParams {
		if strings.Contains(k, "columns") && strings.Contains(k, "orderable") {
			if v[0] == "true" {
				r, _ := regexp.Compile(`\d+`)
				match := r.FindString(k)

				if len(match) > 0 {
					columns = append(columns, qParams["columns["+match+"][name]"][0])
				}
			}
		}
	}

	pageSize, _ := strconv.Atoi(ndu.C.QueryParam("length"))
	offset, _ := strconv.Atoi(ndu.C.QueryParam("start"))
	draw, _ := strconv.Atoi(ndu.C.QueryParam("draw"))

	ndu.draw = draw

	var recordsTotal, recordsFiltered int64

	builder.Count(&recordsTotal)

	ndu.recordsTotal = recordsTotal

	if len(searchStr) > 0 {
		db := builder
		for i, c := range columns {
			if i == 0 {
				db = db.Where(c+" LIKE ?", "%"+searchStr+"%")
			} else {
				db = db.Or(c+` LIKE ?`, "%"+searchStr+"%")
			}
		}

		builder = builder.Where(db)
	}

	builder = builder.Order(orderColumnName + " " + orderDirection)

	builder.Count(&recordsFiltered)

	ndu.recordsFiltered = recordsFiltered

	_ = builder.Offset(offset).Limit(pageSize)

	return ndu
}

func (ndu NewDatatableUtility) AddColumn(colTitle string, callback func(r map[string]interface{}) string) NewDatatableUtility {
	for _, c := range ndu.List.([]map[string]interface{}) {
		nCol := callback(c)

		c[colTitle] = nCol
	}

	return ndu
}

func (ndu NewDatatableUtility) Transform() NewDatatableUtility {
	js, _ := json.Marshal(ndu.List)
	var list []map[string]interface{}
	_ = json.Unmarshal(js, &list)

	ndu.List = list

	return ndu
}

func (ndu NewDatatableUtility) Make() any {
	data := map[string]interface{}{
		"draw":            ndu.draw,
		"recordsTotal":    ndu.recordsTotal,
		"recordsFiltered": ndu.recordsFiltered,
		"data":            ndu.List,
	}

	return data
}
