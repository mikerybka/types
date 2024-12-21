package types

type Schema struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Fields []Field `json:"fields"`
}
