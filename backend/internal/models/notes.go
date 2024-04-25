package models

type Note struct {
	Id        string `json:"id" db:"id"`
	ReagentId string `json:"reagentId,omitempty" db:"reagent_id"`
	Comment   string `json:"comment" db:"comment"`
	Note      string `json:"note" db:"note"`
}

type NoteDTO struct {
	Id        string `json:"id" db:"id"`
	ReagentId string `json:"reagentId,omitempty" db:"reagent_id"`
	Comment   string `json:"comment" db:"comment"`
	Note      string `json:"note" db:"note"`
}

type DeleteNoteDTO struct {
	Id string `json:"id" db:"id"`
}
