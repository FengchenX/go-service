package entity

type Video struct {
	ID   string
	Name string
	Pic  string
	Desc string
	URL  string
}

type FreeVideo struct {
	Video
}

type PaidVideo struct {
	Video
}
