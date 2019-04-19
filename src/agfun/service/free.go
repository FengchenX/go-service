package service

import (
	"agfun/dbcentral/mysqldb"
	"agfun/dto"
	"agfun/entity"
)

func (s *Svc) AddFreeVideos(frees []*entity.FreeVideo) error {
	for _, free := range frees {
		_, i, e := pg.GetFreeVideos(*free, nil)
		if e != nil {
			return e
		}
		if i > 0 {
			continue
		}
		e = pg.AddFreeVideo(free)
		if e != nil {
			return e
		}
	}
	return nil
}
func (s *Svc) GetFreeVideos(req dto.GetVideos) ([]*entity.FreeVideo, error) {
	var id int
	e := s.Dynamic.Get(req.Token, &id)
	if e != nil {
		return nil, e
	}
	var videos []*entity.FreeVideo
	return videos, nil
}
