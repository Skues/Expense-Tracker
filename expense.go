package expense

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type expense struct {
	description string
	amount      float64
	createdAt   time.Time
}

type Expenses []expense

func (e *Expenses) AddExpense(description string, amount float64) {
	var ex expense = expense{
		description,
		amount,
		time.Now()}

	*e = append(*e, ex)
}

func (e *Expenses) UpdateExpense(id int, change string, value []string) error {
	if id < 1 || id-1 > len(*e) {
		return fmt.Errorf("%d out of range.", id)
	}
	for i, _ := range *e {
		if i == id-1 {

		}
		switch change {
		case "description":
			(*e)[i].description = value[0]
		case "amount":
			f, _ := strconv.ParseFloat(value[0], 64)
			(*e)[i].amount = f
		case "both":
			(*e)[i].description = value[0]
			f, _ := strconv.ParseFloat(value[1], 64)
			(*e)[i].amount = f
		}
	}
	return nil
}

func (e *Expenses) ListExpenses() (string, error) {
	if len(*e) == 0 {
		return "", errors.New("List is empty.")
	}
	var sum float64
	var formatted string
	var formattedTime string
	formatted = "\nExpenses:\n----------------\n"
	for _, value := range *e {
		sum += value.amount
		formattedTime = value.createdAt.Format("01 Feb")
		formatted += fmt.Sprintf("\nExpenses:\n----------------\n%s:   %.2f | %s\n------------------\n", value.description, value.amount, formattedTime)
		formatted += fmt.Sprintf("Total Expenses: %.2f", sum)
	}
	return formatted, nil
}

func (e *Expenses) DeleteExpense(id int) error {
	if id < 1 || id-1 > len(*e) {
		return fmt.Errorf("%d out of range.", id)
	}
	index := id - 1
	*e = append((*e)[index+1:], (*e)[:index]...)
	return nil
}
