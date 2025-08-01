package models

type Spending struct {
	Id        string  `json:"id" db:"id"`
	ReagentId string  `json:"reagentId,omitempty" db:"reagent_id"`
	Date      int     `json:"date" db:"date_of_spending"`
	Amount    float64 `json:"amount" db:"amount"`
}

type SpendingDTO struct {
	Id        string  `json:"id" db:"id"`
	ReagentId string  `json:"reagentId" db:"reagent_id" binding:"required"`
	Date      int     `json:"date" db:"date_of_spending" binding:"required,gte=1000000"`
	Amount    float64 `json:"amount" db:"amount" binding:"required,gte=0.000001"`
}

type DeleteSpendingDTO struct {
	Id        string `json:"id" db:"id"`
	ReagentId string `json:"reagentId" db:"reagent_id"`
}
