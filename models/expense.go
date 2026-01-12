package models

import "time"

type Expense struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Category    string    `json:"category"`
}

func NewExpense(id int, description, category string, amount int) Expense {
	return Expense{
		Id:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
		Category:    category,
	}
}
