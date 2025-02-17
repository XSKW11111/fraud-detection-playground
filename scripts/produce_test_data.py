import pandas as pd
import random
from datetime import datetime, timedelta

MERCHANT_NAMES = [
    "Amazon",
    "Walmart",
    "eBay",
    "Target",
    "BestBuy",
]

USER_IDS = [
    "123e4567-e89b-12d3-a456-426614174000",
    "123e4567-e89b-12d3-a456-426614174001", 
    "123e4567-e89b-12d3-a456-426614174002",
    "123e4567-e89b-12d3-a456-426614174003",
    "123e4567-e89b-12d3-a456-426614174004",
]


def produce_test_data(num_transactions: int = 1000, file_path: str = "transaction_test_data.csv"):
    data = []

    for i in range(num_transactions):
        data.append([
            random.choice(USER_IDS),
            (datetime.now() + timedelta(days=random.randint(30, 60), hours=random.randint(0, 23), minutes=random.randint(0, 59))).strftime("%Y-%m-%dT%H:%M:%SZ"),
            random.choice(MERCHANT_NAMES),
            round(random.uniform(1, 100000000) / 100, 3)
        ])

    df = pd.DataFrame(data, columns=["UserId", "Merchant Name", "Amount", "Timestamp"])
    df.sort_values(by="Timestamp", inplace=True)
    df.to_csv(file_path, index=False)

    return file_path, num_transactions


if __name__ == "__main__":
    file_path, num_transactions = produce_test_data(10000, "./data/transaction_test_data.csv")
    print(f"File path: {file_path}")
    print(f"Number of transactions: {num_transactions}")
    print("Finished producing test data")
