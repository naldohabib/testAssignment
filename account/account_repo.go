package account

import (
	"CodeAssignment/model"
)

type AccountRepo interface {
	Insert(data *model.Accounts) (*model.Accounts, error)
	InsertCust(data *model.Customers) (*model.Customers, error)
	FindByAccountNum(id string) (*model.Accounts, error)
	FindByCustomerName(name string) (*model.Customers, error)
	Transfer(id string, data *model.Accounts) (*model.Accounts, error)
}
