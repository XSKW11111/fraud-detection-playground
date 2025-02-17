package main

import (
	"altas.com/fraud/repository"
	"altas.com/fraud/service"
	"altas.com/fraud/utils"
)

func main() {
	db := repository.GetDB()

	repository.InitTables(db)

	transactionService := service.NewTransactionService(db)

	utils.ProcessTransactionCSVFile("./data/transaction_test_data.csv", transactionService)

}
