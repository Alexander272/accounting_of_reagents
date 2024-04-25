package pq_models

import "github.com/lib/pq"

type ReagentDTO struct {
	Id                string         `json:"id" db:"id"`
	Type              string         `json:"type" db:"type"`
	Name              string         `json:"name" db:"name"`
	Uname             string         `json:"uname" db:"uname"`
	Document          string         `json:"document" db:"document"`
	Purity            string         `json:"purity" db:"purity"`
	DateOfManufacture int            `json:"dateOfManufacture" db:"date_of_manufacture"`
	Consignment       string         `json:"consignment" db:"consignment"`
	Manufacturer      string         `json:"manufacturer" db:"manufacturer"`
	ShelfLife         int            `json:"shelfLife" db:"shelf_life"`
	Closet            string         `json:"place_closet" db:"closet"`
	Shelf             int            `json:"place_shelf" db:"shelf"`
	ReceiptDate       int            `json:"incomingControl_receiptData" db:"receipt_date"`
	Amount            string         `json:"incomingControl_amount" db:"amount"`
	ControlDate       int            `json:"incomingControl_date" db:"control_data"`
	Protocol          string         `json:"incomingControl_protocol" db:"protocol"`
	Result            bool           `json:"incomingControl_result" db:"result"`
	Spending          float64        `json:"spending" db:"spending"`
	DateOfExtending   int            `json:"extending_date" db:"date_of_extending"`
	Period            int            `json:"extending_period" db:"period_of_extending"`
	Seizure           string         `json:"seizureInformation" db:"seizure"`
	Disposal          string         `json:"disposalInformation" db:"disposal"`
	Comments          pq.StringArray `json:"comments" db:"comments"`
	Notes             pq.StringArray `json:"notes" db:"notes"`
	Total             int            `db:"total"`
}
