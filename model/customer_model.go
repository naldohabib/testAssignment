package model

type Customers struct {
	CustomerNumber uint   `json:"customer_number,omitempty"`
	Name           string `gorm:"name" json:"name"`
}
