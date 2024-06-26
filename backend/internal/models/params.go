package models

type Params struct {
	Page    *Page
	Sort    []*Sort
	Filters []*Filter
	Search  *Search
	User    *User
}

type Page struct {
	Limit  int
	Offset int
}

type Sort struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type Filter struct {
	Field     string         `json:"field"`
	FieldType string         `json:"fieldType"`
	Values    []*FilterValue `json:"values"`
}
type FilterValue struct {
	CompareType string `json:"compareType"`
	Value       string `json:"value"`
}

type Search struct {
	Value  string
	Fields []string
}
