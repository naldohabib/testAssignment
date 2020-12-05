package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type Accounts struct {
	AccountNumber  uint `json:"account_number,omitempty"`
	CustomerNumber uint `json:"customer_number", omitempty`
	Balance        int  `json:"balance"`
}

type AccountWrap struct {
	Accounts  Accounts  `json:"accounts"`
	Customers Customers `json:"customers"`
}

type AddAccount struct {
	CustomerNumber uint   `json:"customer_number,omitempty"`
	Name           string `gorm:"name" json:"name"`
	Balance        int    `json:"balance"`
}

type FindAccount struct {
	AccountNumber  uint `json:"account_number,omitempty"`
	Name           string `gorm:"name" json:"name"`
	Balance        int    `json:"balance"`
}

// Validate ...
func (u *Accounts) Validate() error {

	if err := validation.Validate(u.Balance, validation.Required); err != nil {
		return errors.New("Balance Cannot be blank")
	}
	return nil
}