package models

type Params struct {
	Sort     []*Sort
	Search   *Search
	Filters  []*NestedFilter
	Page     *Page
	IsPublic bool
	RealmId  string
}

type Page struct {
	Limit  int
	Offset int
}

type Sort struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type Search struct {
	Value  string   `json:"value"`
	Fields []string `json:"fields"`
}

type NestedFilter struct {
	Field  string
	Values []*SingleValue
}

type SingleValue struct {
	Value       string
	CompareType string
}
