package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"altas.com/fraud/model"
	"altas.com/fraud/service"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProcessTransactionCSVFile(filePath string, transactionService *service.TransactionService) {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	fraudTransactions := make([]model.Transaction, 0)

	reader := csv.NewReader(file)

	// Skip header
	if _, err := reader.Read(); err != nil {
		fmt.Println("⚠️ Error reading header:", err)
		return
	}

	// Read each transaction record
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// convert user_id to int64
		userID, err := uuid.Parse(record[0])
		if err != nil {
			fmt.Printf("Error parsing user_id %s: %v\n", record[0], err)
			continue
		}

		timestamp, err := time.Parse(time.RFC3339, record[1])
		if err != nil {
			fmt.Printf("Error parsing timestamp %s: %v\n", record[1], err)
			continue
		}

		// Convert amount to float
		amount, err := decimal.NewFromString(record[3])
		if err != nil {
			fmt.Printf("Error parsing amount %s: %v\n", record[3], err)
			continue
		}

		transaction := model.Transaction{
			ID:           uuid.Must(uuid.NewRandom()),
			UserID:       userID,
			Timestamp:    timestamppb.New(timestamp),
			MerchantName: record[2],
			Amount:       amount,
		}

		err = transactionService.ProcessTransaction(&transaction)
		if err != nil {
			fraudTransactions = append(fraudTransactions, transaction)
		}
	}

	fmt.Println("Fraud transactions:", len(fraudTransactions))

	fmt.Println("Writing fraud transactions to fraud_transactions.csv")
	// Create output CSV file for fraud transactions
	outputFile, err := os.Create("fraud_transactions.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write header
	header := []string{"UserID", "Amount", "MerchantName", "Timestamp"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	// Write fraud transactions
	for _, t := range fraudTransactions {
		record := []string{
			t.UserID.String(),
			t.Amount.String(),
			t.MerchantName,
			t.Timestamp.AsTime().Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing record %v: %v\n", record, err)
			continue
		}
	}

	fmt.Printf("Wrote %d fraud transactions to fraud_transactions.csv\n", len(fraudTransactions))

}
