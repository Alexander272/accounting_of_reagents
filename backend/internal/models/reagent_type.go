package models

type ReagentType struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	RoleId      string `json:"roleId" db:"role_id"`
}

type ReagentTypeDTO struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	RoleId      string `json:"roleId" db:"role_id"`
}

type DeleteReagentTypeDTO struct {
	Id string `json:"id" db:"id"`
}
