package repository

import (
	"CodeAssignment/account"
	"CodeAssignment/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type AccountRepoImpl struct {
	DB *gorm.DB
}

func (a AccountRepoImpl) Transfer(id string, data *model.Accounts) (*model.Accounts, error) {
	err := a.DB.Model(&data).Where("account_number = ?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("AccountRepoImpl.Update Error when query update data with error: %w", err)
	}
	return data, nil
}

func (a AccountRepoImpl) FindByCustomerName(name string) (*model.Customers, error) {
	dataCustomer := new(model.Customers)

	if err := a.DB.Table("customer").Where("customer_number = ?", name).First(&dataCustomer).Error; err != nil {
		fmt.Errorf("[AccountRepoImpl.FindByAccountNum] Error when query get by id with error: %w", err)
		return nil, errors.New("ERROR: Error no data account with id you entered")
	}

	return dataCustomer, nil
}

func (a AccountRepoImpl) FindByAccountNum(id string) (*model.Accounts, error) {
	dataAccount := new(model.Accounts)

	if err := a.DB.Table("account").Where("account_number = ?", id).First(&dataAccount).Error; err != nil {
		fmt.Errorf("[AccountRepoImpl.FindByAccountNum] Error when query get by id with error: %w", err)
		return nil, errors.New("ERROR: Error no data account with id you entered")
	}

	return dataAccount, nil
}

func (a AccountRepoImpl) InsertCust(data *model.Customers) (*model.Customers, error) {

	err := a.DB.Table("customer").Save(&data).Error
	if err != nil {
		fmt.Printf("[AccountRepoImpl.Insert] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data account")
	}

	return data, nil
}

func (a AccountRepoImpl) Insert(data *model.Accounts) (*model.Accounts, error) {

	err := a.DB.Table("account").Save(&data).Error
	if err != nil {
		fmt.Printf("[AccountRepo.Insert] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data account")
	}

	return data, nil
}

func CreateAccountRepo(DB *gorm.DB) account.AccountRepo {
	return &AccountRepoImpl{DB}
}
