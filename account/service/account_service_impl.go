package service

import (
	"CodeAssignment/account"
	"CodeAssignment/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AccountServiceImpl struct {
	accountRepo account.AccountRepo
}

func (a AccountServiceImpl) Transfer(id string, data *model.Accounts) (*model.Accounts, error) {
	_, err := a.accountRepo.FindByAccountNum(id)
	if err != nil {
		return nil, errors.New("userID does not exist")
	}

	user, err := a.Transfer(id, data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a AccountServiceImpl) FindByAccountNum(id string) (*model.Accounts, error) {
	return a.accountRepo.FindByAccountNum(id)
}

func (a AccountServiceImpl) Insert(data *model.AddAccount) (*model.AccountWrap, error) {
	uid, _ := uuid.NewRandom()
	var c = model.Customers{
		CustomerNumber: uint(uid.ID()),
		Name: data.Name,
	}

	ab, _ := a.accountRepo.InsertCust(&c)

	uidd, _ := uuid.NewRandom()
	var account = model.Accounts{
		AccountNumber: uint(uidd.ID()),
		CustomerNumber: ab.CustomerNumber,
		Balance: data.Balance,
	}

	ac, _ := a.accountRepo.Insert(&account)
	var accwrap = model.AccountWrap{
		Accounts: *ac,
		Customers: *ab,
	}

	return &accwrap, nil
}

func CreateAccountService(accountRepo account.AccountRepo) account.AccountService {
	return &AccountServiceImpl{accountRepo}
}
