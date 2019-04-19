package pg

import (
	"agfun/entity"
	"agfun/util"
	"fmt"
)

func GetVideos(video entity.Video, filter *util.PageFilter) ([]*entity.Video, int, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if len(video.ID) > 0 {
		sql = fmt.Sprintf("%s %s id = ?", sql, comma)
		params = append(params, video.ID)
		comma = "AND"
	}
	if len(video.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, video.Name)
		comma = "AND"
	}
	var videos []*entity.Video
	var total int
	db := GetSysDB().Model(&entity.Video{}).Where(sql, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	sql = util.PageFilterSql(sql, "id", filter)
	db = GetSysDB().Where(sql, params...).Find(&videos)
	return videos, total, db.Error
}
