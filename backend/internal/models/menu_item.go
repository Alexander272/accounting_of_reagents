package models

type MenuItem struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsShow      bool   `json:"isShow" db:"is_show"`
	Api         []Api  `json:"api,omitempty"`
}

type MenuItemDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsShow      bool   `json:"isShow"`
}
