package models

type User struct {
	Id   string   `json:"id" db:"id"`
	Name string   `json:"name" db:"name"`
	Role string   `json:"role"`
	Menu []string `json:"menu"`
	// Filters []SIFilter `json:"filters"`

	AccessToken  string `json:"token"`
	RefreshToken string `json:"-"`
}
