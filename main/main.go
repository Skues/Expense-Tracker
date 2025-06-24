package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDesc := addCmd.String("description", "", "Adds the description to an expense")
	addAmount := addCmd.Float64("amount", 0, "Adds the amount to an expense")

	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	viewCategory := viewCmd.String("category", "", "View a specific category of expense")
	viewMonth := viewCmd.String("month", "", "View a specific month of expenses")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", -1, "Deletes a specified ID")
	deleteAll := deleteCmd.Bool("all", false, "Deletes all the expenses")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "Specifies the ID of the expense the user wants to update")
	updateDesc := updateCmd.String("description", "", "Updates the description on a specified expense")
	updateAmount := updateCmd.Float64("amount", 0, "Updades the amount on a specified expense")

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryCategory := summaryCmd.String("category", "", "View the summary of a category of expenses")
	summaryMonth := summaryCmd.String("month", "", "View the summary of a months expenses")

	// view := flag.Bool("view", false, "View all expenses")
	// delete := flag.Int("delete", 0, "Deletes an expense from the list using an ID")
	// update := flag.Int("update", 0, "Updates an expense using its ID")
	// summary := flag.Bool("summary", false, "Displays a summary of the expenses")

	// scanner := bufio.NewScanner(os.Stdin)
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		exp.AddExpense(*addDesc, *addAmount)
		if err := exp.SaveFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "view":
		viewCmd.Parse(os.Args[2:])

		output, err := exp.ListExpenses(*viewMonth, *viewCategory)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, output)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if len(exp) == 0 {
			fmt.Fprint(os.Stderr, "Expense list is empty")
			os.Exit(1)
		}
		if !*deleteAll {
			err := exp.DeleteExpense(*deleteID)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	case "update":
		updateCmd.Parse(os.Args[2:])
		if len(exp) == 0 {
			fmt.Fprint(os.Stderr, "Expense list is empty")
			os.Exit(1)
		}
		exp.UpdateExpense(*updateID, *updateAmount, *updateDesc)
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		output, err := exp.DisplaySummary(*summaryCategory, *summaryMonth)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, output)

	}
	// switch {
	// case *add:
	// 	fmt.Fprintln(os.Stdout, "Enter the name of the Expense (Simple names recommended):")
	// 	scanner.Scan()
	// 	title := scanner.Text()

	// 	fmt.Fprintln(os.Stdout, "Enter the amount without a £ sign:")
	// 	scanner.Scan()
	// 	amount, err := strconv.ParseFloat(scanner.Text(), 64)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// 	exp.AddExpense(title, amount)
	// 	if err := exp.SaveFile(filename); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}

	// case *view:
	// 	output, err := exp.ListExpenses()
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}
	// 	fmt.Fprintln(os.Stdout, output)
	// case *delete != 0:
	// 	if err := exp.DeleteExpense(*delete); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}
	// 	if err := exp.SaveFile(filename); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}

	// case *update != 0:
	// 	var description string
	// 	var amount string
	// 	fmt.Fprintln(os.Stdout, "Enter what you would like to change (description, amount or both):")
	// 	scanner.Scan()
	// 	var change string = scanner.Text()
	// 	switch change {
	// 	case "description":
	// 		fmt.Fprintln(os.Stdout, "Enter the new description:")
	// 		scanner.Scan()

	// 		description = strings.Trim(scanner.Text(), " ")
	// 	case "amount":
	// 		fmt.Fprintln(os.Stdout, "Enter the new expense amount:")
	// 		scanner.Scan()

	// 		amount = strings.Trim(scanner.Text(), " ")

	// 	case "both":
	// 		fmt.Fprintln(os.Stdout, "Enter the new description:")
	// 		scanner.Scan()

	// 		description = strings.Trim(scanner.Text(), " ")
	// 		fmt.Fprintln(os.Stdout, "Enter the new expense amount: ")
	// 		scanner.Scan()

	// 		amount = strings.Trim(scanner.Text(), " ")
	// 	default:
	// 		fmt.Fprintln(os.Stderr, "Specified change is not one of the options.")
	// 		os.Exit(1)

	// 	}
	// 	err := exp.UpdateExpense(*update, change, amount, description)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)

	// 	}
	// 	if err := exp.SaveFile(filename); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}

	// case *summary:
	// 	formatted, err := exp.DisplaySummary()
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Fprintln(os.Stdout, formatted)
	// }

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
