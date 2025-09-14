package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gitshubham45/designPatternGo/splitwise/pkg"
)

func main() {
	fmt.Println("Hello")

	u1 := pkg.NeUser("abhishek", "Abhishek Kumar", "abhishek@gmail.com", "9983345677", 0)
	u2 := pkg.NeUser("rahul", "Rahul Singh", "rahul@gmail.com", "9983345678", 1)
	u3 := pkg.NeUser("shyam", "Shyam Sundar", "shyam@gmail.com", "9983345679", 2)
	u4 := pkg.NeUser("alok", "Alok Singh", "alok@gmail.com", "9983345670", 3)

	groupOne := &pkg.SplitGroup{
		GroupName:  "G1",
		Users:      make([]*pkg.User, 0, 4),
		TotalUsers: 4,
		ExpenseMap: make(map[string]float64),
	}

	groupOne.Users = append(groupOne.Users, u1, u2, u3, u4)

	fmt.Printf("Group one %+v\n", groupOne.Users)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter expense commands (type 'exit' to quit):")

	for {
		fmt.Print(">")
		scanner.Scan()

		input := scanner.Text()

		parts := strings.Fields(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		if strings.ToLower(parts[0]) == "show" {
			if len(parts) < 2 {
				groupOne.ShowAll()

			}
			if len(parts) == 2 {
				groupOne.ShowOne(parts[1])
			}

			continue
			// pkg.show
		}

		expense, err := pkg.ParseExpenseCommand(parts)
		if err != nil {
			fmt.Println("Error : ", err)
			continue
		}

		groupOne.AddExpense(expense)

		fmt.Printf("Parsed Expende %+v\n", expense)
	}
}

/*
 User - userId, name, email, mobile number.
 Expense - id ,amount , paidBy , TYPE( EQUAL, EXACT or PERCENT )
SPLIT Wise - Add


800 equal - 800/4 - 200

u1 - 1200  700

u2 - 800  300

u3 - 150 -> u2 , 350 -> u1. or u3 - > 500 u1
u4 - 150 -> u2 , 350 -> u1 or u4 -> 200 u1 , 300 u2




500 per head

a r s a

EXPENSE abhishek 4 abhishek rahul shyam alok 1000 EQUAL

EXPENSE rahul 4 abhishek rahul shyam alok 1000 EXACT 24 34 56 44

EXPENSE abhishek 4 abhishek rahul shyam alok 1000 EXACT 200 400 100 300

EXPENSE rahul 4 abhishek rahul shyam alok 400 EXACT 400 0 0 0
EXPENSE rahul 4 abhishek rahul shyam alok 800 EXACT 400 200 200 0
EXPENSE rahul 4 abhishek rahul shyam alok 700 EXACT 400 0 100 200
EXPENSE rahul 4 abhishek rahul shyam alok 100 PERCENT 40 10 20 30




*/
