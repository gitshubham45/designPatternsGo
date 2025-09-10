package main

/*
 User - userId, name, email, mobile number.
 Expense - id ,amount , paidBy , TYPE( EQUAL, EXACT or PERCENT )

*/

type expenseType string

const (
	equql   expenseType = "EQUAL"
	exact   expenseType = "EXACT"
	percent expenseType = "PERCENT"
)

type Expense struct {
	ID     string
	Amount float64
	PaidBy string
	Type   expenseType
}
