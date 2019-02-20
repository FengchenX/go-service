package entity

type Video struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pic  string `json:"pic"`
	Desc string `json:"desc"`
	URL  string `json:"url"`
}

type FreeVideo struct {
	Video
}

type PaidVideo struct {
	Video
}
