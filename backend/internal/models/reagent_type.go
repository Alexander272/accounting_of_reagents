package models

type ReagentType struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	RoleId string `json:"roleId" db:"role_id"`
}

type ReagentTypeDTO struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	RoleId string `json:"roleId" db:"role_id"`
}

type DeleteReagentTypeDTO struct {
	Id string `json:"id" db:"id"`
}
