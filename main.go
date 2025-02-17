package main

import (
	"log"

	"altas.com/fraud/repository"
	"altas.com/fraud/service"
	"altas.com/fraud/utils"
)

func main() {
	db := repository.GetDB()

	repository.InitTables(db)

	// For testing purposes, we truncate the tables
	_, err := db.Exec("TRUNCATE TABLE transactions RESTART IDENTITY CASCADE;")

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("TRUNCATE TABLE merchants RESTART IDENTITY CASCADE;")
	if err != nil {
		log.Fatal(err)
	}

	transactionService := service.NewTransactionService(db)

	utils.ProcessTransactionCSVFile("./data/transaction_test_data.csv", transactionService)

}
