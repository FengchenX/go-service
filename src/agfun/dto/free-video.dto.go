package dto

import "agfun/util"

type GetVideos struct {
	Token  string
	Filter *util.PageFilter
}
