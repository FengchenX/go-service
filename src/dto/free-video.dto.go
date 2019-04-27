package dto

import "util"

type GetVideos struct {
	Token  string
	Filter *util.PageFilter
}
