package models

type ReagentList struct {
	Total int        `json:"total" db:"total"`
	List  []*Reagent `json:"list"`
}

type Reagent struct {
	Id                string  `json:"id" db:"id"`
	Type              string  `json:"type" db:"type"`
	Name              string  `json:"name" db:"name"`
	Uname             string  `json:"uname" db:"uname"`
	Document          string  `json:"document" db:"document"`
	Purity            string  `json:"purity" db:"purity"`
	DateOfManufacture int     `json:"dateOfManufacture" db:"date_of_manufacture"`
	Consignment       string  `json:"consignment" db:"consignment"`
	Manufacturer      string  `json:"manufacturer" db:"manufacturer"`
	ShelfLife         int     `json:"shelfLife" db:"shelf_life"`
	Closet            string  `json:"place_closet" db:"closet"`
	Shelf             int     `json:"place_shelf" db:"shelf"`
	ReceiptDate       int     `json:"incomingControl_receiptDate" db:"receipt_date"`
	Amount            string  `json:"incomingControl_amount" db:"amount"`
	ControlDate       int     `json:"incomingControl_date" db:"control_date"`
	Protocol          string  `json:"incomingControl_protocol" db:"protocol"`
	Result            bool    `json:"incomingControl_result" db:"result"`
	Spending          string  `json:"spending" db:"spending"`
	DateOfExtending   int     `json:"extending_date" db:"date_of_extending"`
	Period            int     `json:"extending_period" db:"period_of_extending"`
	Seizure           string  `json:"seizureInformation" db:"seizure"`
	Disposal          string  `json:"disposalInformation" db:"disposal"`
	Comments          string  `json:"comments" db:"comments"`
	Notes             string  `json:"notes" db:"notes"`
	HasRunOut         bool    `json:"hasRunOut" db:"has_run_out"`
	IsOverdue         bool    `json:"isOverdue" db:"is_overdue"`
	SumPeriod         int     `json:"-"`
	ItemStyle         *Styles `json:"itemStyle"`
}

type Styles struct {
	Background string `json:"background,omitempty"`
	TextColor  string `json:"textColor,omitempty"`
}

type EditReagent struct {
	Id                string  `json:"id" db:"id"`
	TypeId            string  `json:"typeId" db:"type_id"`
	Name              string  `json:"name" db:"name"`
	Uname             string  `json:"uname" db:"uname"`
	Document          string  `json:"document" db:"document"`
	Purity            string  `json:"purity" db:"purity"`
	DateOfManufacture int     `json:"dateOfManufacture" db:"date_of_manufacture"`
	Consignment       string  `json:"consignment" db:"consignment"`
	Manufacturer      string  `json:"manufacturer" db:"manufacturer"`
	ShelfLife         int     `json:"shelfLife" db:"shelf_life"`
	Closet            string  `json:"place_closet" db:"closet"`
	Shelf             int     `json:"place_shelf" db:"shelf"`
	ReceiptDate       int     `json:"receiptDate" db:"receipt_date"`
	Amount            float64 `json:"amount" db:"amount"`
	AmountTypeId      string  `json:"amountTypeId" db:"amount_type_id"`
	ControlDate       int     `json:"controlDate" db:"control_date"`
	Protocol          string  `json:"protocol" db:"protocol"`
	Result            bool    `json:"result" db:"result"`
	Seizure           string  `json:"seizure" db:"seizure"`
	Disposal          string  `json:"disposal" db:"disposal"`
}

type ReagentDTO struct {
	Id                string  `json:"id" db:"id"`
	TypeId            string  `json:"typeId" db:"type_id"`
	Name              string  `json:"name" db:"name"`
	Uname             string  `json:"uname" db:"uname"`
	Document          string  `json:"document" db:"document"`
	Purity            string  `json:"purity" db:"purity"`
	DateOfManufacture int     `json:"dateOfManufacture" db:"date_of_manufacture"`
	Consignment       string  `json:"consignment" db:"consignment"`
	Manufacturer      string  `json:"manufacturer" db:"manufacturer"`
	ShelfLife         int     `json:"shelfLife" db:"shelf_life"`
	Closet            string  `json:"place_closet" db:"closet"`
	Shelf             int     `json:"place_shelf" db:"shelf"`
	ReceiptDate       int     `json:"receiptDate" db:"receipt_date"`
	Amount            float64 `json:"amount" db:"amount"`
	AmountTypeId      string  `json:"amountTypeId" db:"amount_type_id"`
	ControlDate       int     `json:"controlDate" db:"control_date"`
	Protocol          string  `json:"protocol" db:"protocol"`
	Result            bool    `json:"result" db:"result"`
	Seizure           string  `json:"seizure" db:"seizure"`
	Disposal          string  `json:"disposal" db:"disposal"`
}

type DeleteReagentDTO struct {
	Id string `json:"id" db:"id"`
}

type ReagentWithRemainder struct {
	Id              string  `json:"id" db:"id"`
	Name            string  `json:"name" db:"name"`
	Document        string  `json:"document" db:"document"`
	Purity          string  `json:"purity" db:"purity"`
	Manufacturer    string  `json:"manufacturer" db:"manufacturer"`
	Amount          float64 `json:"amount" db:"amount"`
	Remainder       float64 `json:"remainder" db:"remainder"`
	HasNotification bool    `json:"hasNotification" db:"has_notification"`
}

type ReagentNotificationDTO struct {
	Id              string `db:"id"`
	HasNotification bool   `db:"has_notification"`
	HasRunOut       bool   `db:"has_run_out"`
}

type ReagentOrderDTO struct {
	List []string `json:"list"`
}
