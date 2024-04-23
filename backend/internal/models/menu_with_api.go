package models

type MenuItemWithApi struct {
	Id             string `db:"id"`
	MenuId         string `db:"menu_item_id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	IsShow         bool   `db:"is_show"`
	ApiId          string `db:"api_id"`
	Method         string `db:"method"`
	Path           string `db:"path"`
	ApiDescription string `db:"api_description"`
}

type MenuWithApiDTO struct {
	Id         string `json:"id"`
	MenuItemId string `json:"menuItemId"`
	ApiId      string `json:"apiId"`
}
