package entity

type Video struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Describe string `json:"describe"`
	Thumb    string `json:"thumb"`
	Creator  string `json:"creator"`
	CreateAt string `json:"create_at"`
}
