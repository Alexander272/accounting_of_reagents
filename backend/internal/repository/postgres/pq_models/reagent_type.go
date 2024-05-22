package pq_models

import "github.com/lib/pq"

type ReagentTypeDTO struct {
	Id          string         `json:"id" db:"id"`
	Name        string         `json:"name" db:"name"`
	Number      int            `json:"number" db:"number"`
	Description string         `json:"description" db:"description"`
	RoleId      pq.StringArray `json:"roleId" db:"role_id"`
}
