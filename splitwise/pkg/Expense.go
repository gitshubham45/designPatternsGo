package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

/*
 User - userId, name, email, mobile number.
 Expense - id ,amount , paidBy , TYPE( EQUAL, EXACT or PERCENT )

*/

type Expense struct {
	ID     string
	Amount float64
	PaidBy string
	Split  string
	Values []float64
	Users  []string
}

func ParseExpenseCommand(parts []string) (*Expense, error) {

	fmt.Println("Parts : ", parts)

	if len(parts) < 5 || strings.ToUpper(parts[0]) != "EXPENSE" {
		return nil, fmt.Errorf("invlaid command format")
	}

	paidBy := parts[1]

	numUsers, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invlaid number of users")
	}

	users := parts[3 : 3+numUsers]

	amount, err := strconv.ParseFloat(parts[3+numUsers], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %f", amount)
	}

	splitType := strings.ToUpper(parts[4+numUsers])

	// values (optional) depending on the split type
	values := []float64{}
	if splitType == "EXACT" || splitType == "PERCENT" {
		for _, v := range parts[5+numUsers:] {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("invlaid value: %s", v)
			}
			if splitType == "EXACT" {
				values = append(values, f)
			} else {
				values = append(values, amount*f/100)
			}
		}
	} else if splitType == "EQUAL" {
		for i := 0; i < numUsers; i++ {
			values = append(values, amount/float64(numUsers))
		}
	}

	expense := &Expense{
		ID:     uuid.NewString(),
		PaidBy: paidBy,
		Users:  users,
		Split:  splitType,
		Values: values,
		Amount: amount,
	}
	return expense, nil
}

// EXPENSE Abhishek 4 Abhishek Rahul Shyam Alok EQUAL 1000
