package account

import (
	"CodeAssignment/model"
)

type AccountService interface {
	Insert(data *model.AddAccount) (*model.AccountWrap, error)
	FindByAccountNum(id string) (*model.Accounts, error)
	//FindAllCustomer() (*[]model.Customers, error)
	Transfer(id string, data *model.Accounts) (*model.Accounts, error)
}
