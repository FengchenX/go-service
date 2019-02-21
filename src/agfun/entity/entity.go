package entity

type Video struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pic  string `json:"pic"`
	Desc string `json:"desc"`
	URL  string `json:"url"`
}

type FreeVideo struct {
	ID      string `json:"id"`
	VideoID string `json:"video_id"`
}

type PaidVideo struct {
	ID      string `json:"id"`
	VideoID string `json:"video_id"`
}
