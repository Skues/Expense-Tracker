package expense

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	Description string
	Amount      float64
	CreatedAt   time.Time
}

type Expenses []Expense

type Personal struct {
	Balance float64
	Salary  float64
}

func (e *Expenses) AddExpense(Description string, Amount float64) {
	var ex Expense = Expense{
		Description,
		Amount,
		time.Now()}

	*e = append(*e, ex)
	AlterBalance(Amount)
}

func (e *Expenses) UpdateExpense(id int, amount float64, description string) error {
	if id < 1 || id-1 > len(*e) {
		return fmt.Errorf("%d out of range.", id)
	}

	for i, _ := range *e {
		if i == id-1 {
			if amount != 0 && description != "" {
				(*e)[i].Amount = amount
				(*e)[i].Description = description
			} else if amount != 0 {
				(*e)[i].Amount = amount
			} else if description != "" {
				(*e)[i].Description = description
			}
		}

	}
	fmt.Println(*e)
	return nil
}

func (e *Expenses) ListExpenses(month string, category string) (string, error) {
	if len(*e) == 0 {
		return "", errors.New("List is empty.")
	}
	var formatted string
	var formattedTime string
	var counter int

	formatted = "\nExpenses:\n"
	if month == "" && category == "" {
		for index, value := range *e {
			formattedTime = value.CreatedAt.Format("02 Jan")
			formatted += fmt.Sprintf("\n%d - %s:   £%.2f | %s\n------------------\n", index+1, value.Description, value.Amount, formattedTime)
		}
	} else if month != "" && category == "" {
		formatted += fmt.Sprintf("----%s----", month)
		for index, value := range *e {
			formattedTime = value.CreatedAt.Format("02 Jan")
			if value.CreatedAt.Month().String() == month {
				counter += 1
				formatted += fmt.Sprintf("\n%d - %s:   £%.2f | %s\n------------------\n", index+1, value.Description, value.Amount, formattedTime)
			}
		}
		if counter == 0 {
			formatted = fmt.Sprintf("No expenses found for %s.", month)
		}
	}

	return formatted, nil
}

func (e *Expenses) DisplaySummary(category string, month string) (string, error) {
	var sum float64
	if len(*e) <= 0 {
		return "", errors.New("No items in the list of expenses.")
	}
	if category == "" && month == "" {
		for _, value := range *e {
			sum += value.Amount
		}
	} else if month != "" {
		for _, value := range *e {
			if value.CreatedAt.Month().String() == month {
				sum += value.Amount
			}
		}
	}

	formatted := fmt.Sprintf("Your total expenses are: £%.2f", sum)
	return formatted, nil
}

func (e *Expenses) DeleteExpense(id int) error {
	if id < 1 || id-1 > len(*e) {
		return fmt.Errorf("%d out of range.", id)
	}
	index := id - 1
	AlterBalance(-((*e)[index].Amount))
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

func (e *Expenses) SetBalance(balance float64) error {
	file, err := os.Create("balance.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	balanceStr := strconv.FormatFloat(balance, 'f', -1, 64)
	_, err = file.Write([]byte(balanceStr))
	if err != nil {
		return err
	}
	file.Sync()
	return nil
}

func AlterBalance(amount float64) {
	balance, _ := CurrentBalance()
	newBalance := balance - amount
	saveBalance(newBalance)
}

func CurrentBalance() (float64, error) {
	bytes, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, err
	}
	balance, err := strconv.ParseFloat(string(bytes), 64)
	return balance, err

}

func saveBalance(balance float64) {
	file, err := os.Create("balance.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()
	_, err = file.WriteString(strconv.FormatFloat(balance, 'f', -1, 64))
	if err != nil {
		fmt.Fprint(os.Stderr, "testsing")
		fmt.Fprintln(os.Stderr, err)
	}
}
