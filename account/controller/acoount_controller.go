package controller

import (
	"CodeAssignment/account"
	"CodeAssignment/model"
	"CodeAssignment/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	accountService account.AccountService
}

func (h AccountHandler) insert(resp http.ResponseWriter, req *http.Request) {
	user := new(model.AddAccount)
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[AccountHandler.insert] Error when decoder data from body with error : %v\n", err)
		return
	}

	response, err := h.accountService.Insert(user)
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, err.Error())
		fmt.Printf("[AccountHandler.insertData] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.HandleSuccess(resp, http.StatusOK, response)

}

func (h AccountHandler) findAccountByNum(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id := pathVar["account_number"]


	data, err := h.accountService.FindByAccountNum(id)
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, "User ID not Found!")
		fmt.Printf("[UserHandler.getByID]Error when request data with error : %v \n", err)
		return
	}

	//var accName model.Customers
	//
	//account := model.FindAccount{
	//	AccountNumber: data.AccountNumber,
	//	Name: accName.Name,
	//	Balance: data.Balance,
	//}

	utils.HandleSuccess(resp, http.StatusOK, data)
}

func (h AccountHandler) transfer(resp http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)
	id := pathVar["account_number"]

	_, err := h.accountService.FindByAccountNum(id)
	if err != nil {
		fmt.Printf("[AccountHandler.trandfer] Error when check id to usecase with error: %v\n", err)
		utils.HandleError(resp, http.StatusBadRequest, "ID DOES NOT EXIST")
		return
	}

	var transfer = model.Accounts{}

	err = json.NewDecoder(req.Body).Decode(&transfer)
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, "Oopss, something error")
		fmt.Printf("[AccountHandler.getData] Error when decode data with error: %v\n", err)
		return
	}

	dataUpdate, err := h.accountService.Transfer(id, &transfer)
	if err != nil {
		utils.HandleError(resp, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.update] Error when send data to usecase with error : %v", err)
		return
	}

	utils.HandleSuccess(resp, http.StatusOK, dataUpdate)
}

func CreateAccountHandler(r *mux.Router, accountService account.AccountService)  {
	accHandler := AccountHandler{accountService}

	r.HandleFunc("/account", accHandler.insert).Methods(http.MethodPost)
	r.HandleFunc("/account/{account_number}", accHandler.findAccountByNum).Methods(http.MethodGet)
	r.HandleFunc("/account/{from_account_number}/transfer", accHandler.transfer).Methods(http.MethodPut)

}
