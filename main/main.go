package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"project.com/expense"
)

func main() {

	filename := "expenses.csv"
	exp, err := LoadCSV(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	add := flag.Bool("add", false, "Adds an expense to the list")
	view := flag.Bool("view", false, "View all expenses")
	delete := flag.Int("delete", 0, "Deletes an expense from the list using an ID")
	update := flag.Int("update", 0, "Updates an expense using its ID")
	summary := flag.Bool("summary", false, "Displays a summary of the expenses")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	switch {
	case *add:
		fmt.Fprintln(os.Stdout, "Enter the name of the Expense (Simple names recommended):")
		scanner.Scan()
		title := scanner.Text()

		fmt.Fprintln(os.Stdout, "Enter the amount without a £ sign:")
		scanner.Scan()
		amount, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		exp.AddExpense(title, amount)
		if err := exp.SaveFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

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
		if err := exp.SaveFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *update != 0:
		var description string
		var amount string
		fmt.Fprintln(os.Stdout, "Enter what you would like to change (description, amount or both):")
		scanner.Scan()
		var change string = scanner.Text()
		switch change {
		case "description":
			fmt.Fprintln(os.Stdout, "Enter the new description:")
			scanner.Scan()

			description = strings.Trim(scanner.Text(), " ")
		case "amount":
			fmt.Fprintln(os.Stdout, "Enter the new expense amount:")
			scanner.Scan()

			amount = strings.Trim(scanner.Text(), " ")

		case "both":
			fmt.Fprintln(os.Stdout, "Enter the new description:")
			scanner.Scan()

			description = strings.Trim(scanner.Text(), " ")
			fmt.Fprintln(os.Stdout, "Enter the new expense amount: ")
			scanner.Scan()

			amount = strings.Trim(scanner.Text(), " ")
		default:
			fmt.Fprintln(os.Stderr, "Specified change is not one of the options.")
			os.Exit(1)

		}
		err := exp.UpdateExpense(*update, change, amount, description)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

		}
		if err := exp.SaveFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
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

func LoadCSV(filename string) (expense.Expenses, error) {
	var exp expense.Expenses
	file, err := os.Open(filename)
	if err != nil {
		return expense.Expenses{}, nil
	}

	records, err := csv.NewReader(file).ReadAll()

	for _, record := range records {
		desc := record[0]
		amo, _ := strconv.ParseFloat(record[1], 64)
		time, _ := time.Parse("02-01-06 15:04", record[2])
		data := expense.Expense{
			Description: desc,
			Amount:      amo,
			CreatedAt:   time,
		}
		exp = append(exp, data)
	}
	return exp, nil

}
