package repository

import "database/sql"

func InitTables(db *sql.DB) {
	userRepository := NewUserRepository(db)
	userRepository.initUserTable()

	transactionRepository := NewTransactionRepository(db)
	transactionRepository.initTransactionTable()

	merchantRepository := NewMerchantRepository(db)
	merchantRepository.initMerchantTable()

}
