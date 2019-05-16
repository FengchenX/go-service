package dto

type Search []Row

type Row struct {
	Name    string      `json:"name"`
	Operate string      `json:"operate"`
	Value   interface{} `json:"value"`
}
