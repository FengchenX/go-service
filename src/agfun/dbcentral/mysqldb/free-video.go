package mysqldb

import (
	"agfun/entity"
	"agfun/util"
	"fmt"
)

func AddFreeVideo(free *entity.FreeVideo) error {
	db := GetSysDB().Create(free)
	return db.Error
}
func GetFreeVideos(free entity.FreeVideo, filter *util.PageFilter) ([]*entity.FreeVideo, int, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if len(free.ID) > 0 {
		sql = fmt.Sprintf("%s %s uid = ?", sql, comma)
		params = append(params, free.ID)
		comma = "AND"
	}
	if len(free.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, free.Name)
		comma = "AND"
	}
	var frees []*entity.FreeVideo
	var total int
	db := GetSysDB().Model(&entity.FreeVideo{}).Where(sql, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	sql = util.PageFilterSql(sql, "id", filter)
	db = GetSysDB().Where(sql, params...).Find(&frees)
	return frees, total, db.Error
}
