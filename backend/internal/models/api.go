package models

type Api struct {
	Id          string `json:"id" db:"id"`
	Path        string `json:"path" db:"path"`
	Method      string `json:"method" db:"method"`
	Description string `json:"description" db:"description"`
}

type ApiDTO struct {
	Id          string `json:"id" db:"id"`
	Path        string `json:"path" db:"path"`
	Method      string `json:"method" db:"method"`
	Description string `json:"description" db:"description"`
}
