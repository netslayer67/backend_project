package transactionsdto

import "time"

type TransactionResponse struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    bool      `json:"category"`
}
