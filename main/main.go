package main

import (
	"flag"

	"project.com/expense"
)

func main() {
	exp := expense.Expenses{}
	add := flag.String("add", "", "Adds an expense to the list")
	view := flag.Bool("view", false, "View all expenses")
	delete := flag.Int("delete", 0, "Deletes an expense from the list using an ID")
	update := flag.Int("update", 0, "Updates an expense using its ID")
}
