package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

// page filter
// swagger:model
type PageFilter struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

func ParsePageFilter(c iris.Context) *PageFilter {
	filter := PageFilter{
		c.URLParam("sort"),
		c.URLParam("order"),
		c.URLParam("limit"),
		c.URLParam("offset"),
	}
	if len(filter.Sort) == 0 && len(filter.Order) == 0 && len(filter.Limit) == 0 && len(filter.Offset) == 0 {
		return nil
	}
	return &filter
}

func PageFilterSql(sql *gorm.DB, keyCol string, filter *PageFilter) *gorm.DB {
	if sql ==nil || filter == nil {
		return sql
	}
	if len(filter.Sort) > 0 && len(filter.Order) == 0 {
		filter.Order = "DESC"
	}
	if len(filter.Sort) > 0 && filter.Sort != keyCol {
		sql = sql.Order(fmt.Sprintf("%s %s", filter.Sort, filter.Order))
		if len(keyCol) > 0 {
			sql = sql.Order(keyCol)
		}
	} else if len(keyCol) > 0 {
		sql = sql.Order(fmt.Sprintf("%s %s", keyCol, filter.Order))
	}
	if len(filter.Offset) > 0 && len(filter.Limit) > 0 {
		sql = sql.Limit(filter.Limit).Offset(filter.Offset)
	}
	return sql
}
