package models

import "time"

type Expense struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

func NewExpense(id int, description string, amount int) Expense {
	return Expense{
		Id:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
}
