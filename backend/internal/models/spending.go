package models

type Spending struct {
	Id        string  `json:"id" db:"id"`
	ReagentId string  `json:"reagentId,omitempty" db:"reagent_id"`
	Date      int     `json:"date" db:"date_of_spending"`
	Amount    float64 `json:"amount" db:"amount"`
}

type SpendingDTO struct {
	Id        string  `json:"id" db:"id"`
	ReagentId string  `json:"reagentId" db:"reagent_id"`
	Date      int     `json:"date" db:"date_of_spending"`
	Amount    float64 `json:"amount" db:"amount"`
}

type DeleteSpendingDTO struct {
	Id string `json:"id" db:"id"`
}
