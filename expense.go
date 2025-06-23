package expense

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Description string
	Amount      float64
	CreatedAt   time.Time
}

type Expenses []Expense

func (e *Expenses) AddExpense(Description string, Amount float64) {
	var ex Expense = Expense{
		Description,
		Amount,
		time.Now()}

	*e = append(*e, ex)
}

func (e *Expenses) UpdateExpense(id int, change string, amount string, description string) error {
	if id < 1 || id-1 > len(*e) {
		return fmt.Errorf("%d out of range.", id)
	}
	for i, _ := range *e {
		if i == id-1 {
			switch change {
			case "description":
				(*e)[i].Description = description
			case "amount":
				f, _ := strconv.ParseFloat(strings.Replace(amount, " ", "", -1), 64)
				(*e)[i].Amount = f
			case "both":
				(*e)[i].Description = description
				f, _ := strconv.ParseFloat(amount, 64)
				(*e)[i].Amount = f
			}
		}

	}
	fmt.Println(*e)
	return nil
}

func (e *Expenses) ListExpenses() (string, error) {
	if len(*e) == 0 {
		return "", errors.New("List is empty.")
	}
	var sum float64
	var formatted string
	var formattedTime string
	formatted = "\nExpenses:\n----------------"
	for index, value := range *e {
		sum += value.Amount
		formattedTime = value.CreatedAt.Format("02 Jan")
		formatted += fmt.Sprintf("\n%d - %s:   £%.2f | %s\n------------------\n", index+1, value.Description, value.Amount, formattedTime)
	}
	return formatted, nil
}

func (e *Expenses) DisplaySummary() (string, error) {
	var sum float64
	if len(*e) <= 0 {
		return "", errors.New("No items in the list of expenses.")
	}
	for _, value := range *e {
		sum += value.Amount
	}
	formatted := fmt.Sprintf("Your total expenses are: £%.2f", sum)
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

func (e *Expenses) SaveFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, value := range *e {
		amountStr := strconv.FormatFloat(value.Amount, 'f', 2, 64)
		timeStr := value.CreatedAt.Format("02-01-06 15:04")
		row := []string{value.Description, amountStr, timeStr}
		err := csvWriter.Write(row)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil

}
