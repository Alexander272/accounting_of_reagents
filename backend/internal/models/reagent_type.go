package models

type ReagentType struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	IsPublic    bool   `json:"isPublic" db:"is_public"`
}

type GetReagentTypeDTO struct {
	IsPublic bool `json:"isPublic" db:"is_public"`
}

type ReagentTypeDTO struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	IsPublic    bool   `json:"isPublic" db:"is_public"`
}

type DeleteReagentTypeDTO struct {
	Id string `json:"id" db:"id"`
}
