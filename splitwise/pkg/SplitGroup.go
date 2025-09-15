package pkg

import (
	"fmt"
	"sync"
)

type SplitGroup struct {
	GroupName  string
	Users      []*User
	Expenses   []*Expense
	ExpenseMap map[string]float64
	TotalUsers int
	mu         sync.Mutex
}

func (s *SplitGroup) ShowOne(currentUserName string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("Showing for user %s\n", currentUserName)

	for _, u := range s.Users {
		if currentUserName == u.UserName {
			continue
		}

		A_paid_for_B := fmt.Sprintf("%s-%s", currentUserName, u.UserName)
		B_paid_for_A := fmt.Sprintf("%s-%s", u.UserName, currentUserName)

		amountAB := s.ExpenseMap[A_paid_for_B]
		amountBA := s.ExpenseMap[B_paid_for_A]

		// Ideally, only one of these should be > 0
		if amountAB > 0 {
			fmt.Printf("%s owes %s : %.2f\n", u.UserName, currentUserName, amountAB)
		} else if amountBA > 0 {
			fmt.Printf("%s owes %s : %.2f\n", currentUserName, u.UserName, amountBA)
		}
	}

	fmt.Println("------------------------------")
}

func (s *SplitGroup) ShowAll() {

	fmt.Println("Show all")

	for _, u := range s.Users {
		s.ShowOne(u.UserName)
	}
}

func (s *SplitGroup) AddExpense(expense *Expense) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Expenses = append(s.Expenses, expense)

	// string format - userName-paidBy
	for _, u := range s.Users {
		if expense.PaidBy == u.UserName {
			continue
		}

		amountPaid := expense.Values[u.Index]

		fmt.Println("index", u.Index)

		fmt.Println(expense.Values)

		A_paid_for_B_payment_mapping_string := fmt.Sprintf("%s-%s", expense.PaidBy, u.UserName)
		B_paid_for_A_payment_mapping_string := fmt.Sprintf("%s-%s", u.UserName, expense.PaidBy)

		// fmt.Println("Amount : ", amountPaid)

		// A is paying for B - amount P  ( AB) > 0

		// 1. if B has paid Q, for A in past (BA) > 0

		// total paid By A for B - ( if P > Q) then BA = 0 and AB = P - BA

		// if (Q > P) then AB = 0 , BA = BA - P
		if prev := s.ExpenseMap[B_paid_for_A_payment_mapping_string]; prev > 0 {
			if prev > amountPaid {
				s.ExpenseMap[B_paid_for_A_payment_mapping_string] = prev - amountPaid
				s.ExpenseMap[A_paid_for_B_payment_mapping_string] = 0
			} else {
				s.ExpenseMap[B_paid_for_A_payment_mapping_string] = 0
				s.ExpenseMap[A_paid_for_B_payment_mapping_string] += amountPaid - prev
			}
		} else {
			// no reverse entry, simply add to A-B
			s.ExpenseMap[A_paid_for_B_payment_mapping_string] += amountPaid
		}
	}
}
