package models

type MenuItem struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Method      string `json:"method" db:"method"`
	Description string `json:"description" db:"description"`
	IsShow      bool   `json:"isShow" db:"is_show"`
}

type MenuItemDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Method      string `json:"method"`
	Description string `json:"description"`
	IsShow      bool   `json:"isShow"`
}
