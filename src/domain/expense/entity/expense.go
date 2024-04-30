package expense

import (
	"errors"
	"time"

	"github.com/IsaqueAmorim/money-wise/src/domain/shared/category"
	"github.com/google/uuid"
)

type expense struct {
	ID           uuid.UUID
	Title        string
	Description  string
	Categories   []category.Category
	Ammount      float64
	Installments []installment
	Paid         bool
	CreatedAt    time.Time
	Date         time.Time
}

type installment struct {
	ID             uuid.UUID
	ExpenseID      uuid.UUID
	InstallmentNum int
	DueDate        time.Time
	PaidDate       time.Time
	Value          float64
}

func (e *expense) NewInstallment(numberOfInstallments int, dueDate time.Time) {

	for i := 0; i <= numberOfInstallments; i++ {
		e.Installments[i] = installment{
			ID:             uuid.New(),
			ExpenseID:      e.ID,
			InstallmentNum: i,
			DueDate:        dueDate.Local().AddDate(0, i, 0),
			PaidDate:       time.Time{},
		}
	}
}

func NewExpense(title, description string, categories []category.Category, ammount float64, dueDate time.Time, sendNotification, paid bool, installments int) *expense {

	expense := &expense{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Categories:  categories,
		Ammount:     ammount,
		CreatedAt:   time.Now(),
		Paid:        paid,
	}
	expense.NewInstallment(installments, dueDate)

	return expense
}

func (e *expense) PayExpense() error {

	for i := range e.Installments {
		if e.Installments[i].PaidDate != *new(time.Time) {
			e.Installments[i].PaidDate = time.Now()
		}
	}

	return e.pay()
}

func (e *expense) PayInstallment(installmentNum int) error {

	LAST_ITEM := len(e.Installments) - 1

	if len(e.Installments) < installmentNum {

		e.Installments[installmentNum].PaidDate = time.Now()
		if installmentNum == LAST_ITEM {
			e.pay()
		}
		return nil
	}
	return nil
}

func (e expense) pay() error {

	if e.Paid {
		return errors.New("The invoice has already been paid")
	}
	e.pay()

	return nil
}

func (e *expense) ChangeTitle(title string) {
	e.Title = title
}

func (e *expense) ChangeDescription(description string) {
	e.Description = description
}

func (e *expense) ChangeAmmount(ammount float64) {
	e.Ammount = ammount
}

func (e *expense) ChangeCategories(categories []category.Category) {
	e.Categories = categories
}
