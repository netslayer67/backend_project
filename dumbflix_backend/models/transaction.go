package models

import "time"

type Transaction struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    bool      `json:"category"`
}
