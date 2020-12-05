package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	accountHandler "CodeAssignment/account/controller"
	accountRepo "CodeAssignment/account/repository"
	accountService "CodeAssignment/account/service"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	conStr := viper.GetString("database.postgres")
	port := viper.GetString("port.port")

	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error())
	}

	defer db.Close()

	//db.Debug().AutoMigrate(
	//	&model.Accounts{},
	//	&model.Customers{},
	//)
	//db.Model(&model.Customers{}).AddForeignKey("customer_number", "accounts(account_number)", "CASCADE", "CASCADE")

	router := mux.NewRouter().StrictSlash(true)

	accountRepo := accountRepo.CreateAccountRepo(db)
	accountService := accountService.CreateAccountService(accountRepo)
	accountHandler.CreateAccountHandler(router, accountService)

	fmt.Println("Starting web server at port : ", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal()
	}
}
