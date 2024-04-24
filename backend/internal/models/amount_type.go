package models

type AmountType struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Number int    `json:"number" db:"number"`
}

type AmountTypeDTO struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Number int    `json:"number" db:"number"`
}

type DeleteAmountTypeDTO struct {
	Id string `json:"id" db:"id"`
}
