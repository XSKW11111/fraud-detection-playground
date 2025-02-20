# fraud-detection-playground

## How to run

### Run the script to produce test data

```bash
python3 scripts/produce_test_data.py
```

### start postgres

```bash
docker pull postgres
docker run --name postgres-container -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=transactions_db -p 5432:5432 -d postgres
```

### Run the main program

```bash
go run main.go
```


## fraud detection rules

### Rule 1: If the amount is 5x greater than or equal to user average tranactions, it is a fraud transaction
    - user average transactions is the average amount of transactions for a user
    - user receive large amount of money a lot more than the user normal transactions, it is highly likely a fraud transaction

### Rule 2: if the user has more than 5 transactions in 5 minutes, it is a fraud transaction
    - user receive large batch of transactions in a short period of time, it is highly likely a fraud transaction

### Rule 3: if user has large first time transaction, it is a fraud transaction
    - user receive large amount of money for the first time, it is highly likely a fraud transaction

### Rule 4: if the merchant name is in hight risk list, it is a fraud transaction or if the merchant receive large amount for the first time, it is a fraud transaction
    - merchant name is in hight risk list, it is a fraud transaction
    - merchant receive large amount of money for the first time, it is a fraud transaction

### Rule 5: if the merchant receive large amount that are 5 times of the average amount, it is a fraud transaction
    - merchant receive large batch of transactions that are 5 times of the average amount, it is highly likely a fraud transaction




