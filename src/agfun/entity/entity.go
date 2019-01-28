package entity

type Video struct {
	ID   string
	Name string
	Desc string
	URL  string
}

type FreeVideo struct {
	Video
}

type PaidVideo struct {
	Video
}
