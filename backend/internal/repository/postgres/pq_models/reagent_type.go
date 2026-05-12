package pq_models

type ReagentTypeDTO struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	IsPublic    bool   `json:"isPublic" db:"is_public"`
}
