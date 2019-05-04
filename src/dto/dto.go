package dto

import "entity"

type Search []Row

type Row struct {
	Name    string `json:"name"`
	Operate string `json:"operate"`
	Value   interface{} `json:"value"`
}

type Video struct {
	entity.Video
}

type Videos struct {
	Total int `json:"total"`
	Videos []Video `json:"videos"`
}
