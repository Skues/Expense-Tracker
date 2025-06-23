package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"project.com/expense"
)

func main() {

	exp := expense.Expenses{}

	add := flag.Bool("add", false, "Adds an expense to the list")
	view := flag.Bool("view", false, "View all expenses")
	delete := flag.Int("delete", 0, "Deletes an expense from the list using an ID")
	update := flag.Int("update", 0, "Updates an expense using its ID")
	summary := flag.Bool("summary", false, "Displays a summary of the expenses")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	switch {
	case *add:
		fmt.Println("Testing")
		fmt.Fprintln(os.Stdout, "Enter the name of the Expense (Simple names recommended): ")
		scanner.Scan()
		title := scanner.Text()

		fmt.Fprintln(os.Stdout, "Enter the amount without a £ sign: ")
		scanner.Scan()
		amount, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		exp.AddExpense(title, amount)
	case *view:
		output, err := exp.ListExpenses()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Fprintln(os.Stdout, output)
	case *delete != 0:
		if err := exp.DeleteExpense(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case *update != 0:
		var description string
		var amount string
		fmt.Fprintln(os.Stdout, "Enter what you would like to change (description, expense or both): ")
		scanner.Scan()
		var change string = scanner.Text()
		switch change {
		case "description":
			fmt.Fprintln(os.Stdout, "Enter the new description: ")
			scanner.Scan()

			description = scanner.Text()
		case "expense":
			fmt.Fprintln(os.Stdout, "Enter the new expense amount: ")
			scanner.Scan()

			amount = scanner.Text()

		case "both":
			fmt.Fprintln(os.Stdout, "Enter the new description: ")
			scanner.Scan()

			description = scanner.Text()
			fmt.Fprintln(os.Stdout, "Enter the new expense amount: ")
			scanner.Scan()

			amount = scanner.Text()
		default:
			fmt.Fprintln(os.Stderr, "Specified change is not one of the options.")
			os.Exit(1)

		}
		err := exp.UpdateExpense(*update, change, description, amount)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

		}
	case *summary:
		formatted, err := exp.DisplaySummary()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, formatted)
	}

}
