package models

type Extending struct {
	Id        string `json:"id" db:"id"`
	ReagentId string `json:"reagentId,omitempty" db:"reagent_id"`
	Date      int    `json:"date" db:"date_of_extending"`
	Period    int    `json:"period" db:"period_of_extending"`
}

type ExtendingDTO struct {
	Id        string `json:"id" db:"id"`
	ReagentId string `json:"reagentId" db:"reagent_id"`
	Date      int    `json:"date" db:"date_of_extending"`
	Period    int    `json:"period" db:"period_of_extending"`
}

type DeleteExtendingDTO struct {
	Id string `json:"id" db:"id"`
}
