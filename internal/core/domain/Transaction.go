package domain

import "time"

type Transaction struct {
	Id                 int    `json:"id"`
	TransactionType    int    `json:"transaction_type"`
	SourceAccount      int    `json:"source_account"`
	DestinationAccount int    `json:"destination_account"`
	Status             string `json:"status"`
	Amount             int    `json:"amount"`
	TransactionDate    time.Time
}
