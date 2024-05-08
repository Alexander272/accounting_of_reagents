package models

type AmountType struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Number      int    `json:"number" db:"number"`
}

type AmountTypeDTO struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Number      int    `json:"number" db:"number"`
}

type DeleteAmountTypeDTO struct {
	Id string `json:"id" db:"id"`
}
