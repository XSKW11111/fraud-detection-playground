package repository

import (
	"database/sql"
	"log"
)

func InitTables(db *sql.DB) {
	userRepository := NewUserRepository(db)
	err := userRepository.initUserTable()
	if err != nil {
		log.Println("Error initializing user table", err)
	}

	transactionRepository := NewTransactionRepository(db)
	err = transactionRepository.initTransactionTable()
	if err != nil {
		log.Println("Error initializing transaction table", err)
	}
	if err != nil {
		log.Println("Error initializing transaction table", err)
	}

	merchantRepository := NewMerchantRepository(db)
	err = merchantRepository.initMerchantTable()
	if err != nil {
		log.Println("Error initializing merchant table", err)
	}

}
