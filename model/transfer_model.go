package model

type Transfer struct {
	ToAccountNumber uint `json:"to_account_number"`
	Amount int `json:"amount"`
}
