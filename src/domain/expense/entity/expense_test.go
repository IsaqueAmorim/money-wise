package expense

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var data = &expense{
	ID:          uuid.New(),
	Title:       "Title Test",
	Description: "Description Test",
	Ammount:     100.00,
}

func TestGivenValidParameters_WhenCreatingExpense_CreateExpenseHasExpectedValues(t *testing.T) {
	expense := NewExpense(data.Title, data.Description, data.Categories, data.Ammount, data.Date, data.Paid, 0)

	assert.Equal(t, expense.Title, data.Title)
	assert.Equal(t, expense.Description, data.Description)
	assert.Equal(t, expense.Ammount, data.Ammount)
	assert.NotEqual(t, expense.CreatedAt, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, len(expense.Installments), 0)

}

func TestGivenInstallmentNum_WhenCreateInstallmentsWithDueDate_CreateExpectedIstallmentsNumAndDates(t *testing.T) {
	expense := NewExpense(data.Title, data.Description, data.Categories, data.Ammount, time.Now(), data.Paid, 4)

	assert.Equal(t, len(expense.Installments), 4)
	assert.Equal(t, expense.Installments[0].InstallmentNum, 0)
	assert.Equal(t, expense.Installments[1].InstallmentNum, 1)
	assert.Equal(t, expense.Installments[2].InstallmentNum, 2)
	assert.Equal(t, expense.Installments[3].InstallmentNum, 3)
	assert.Equal(t, expense.Installments[1].DueDate, expense.Date.AddDate(0, 1, 0))
	assert.Equal(t, expense.Installments[2].DueDate, expense.Date.AddDate(0, 2, 0))
	assert.Equal(t, expense.Installments[3].DueDate, expense.Date.AddDate(0, 3, 0))
}

func TestGivenIntallmentNum_WhenPayInstallmentByNum_ThenPayInstallmentSucess(t *testing.T) {
	expense := NewExpense(data.Title, data.Description, data.Categories, data.Ammount, time.Now(), data.Paid, 4)
	TimeDifference := expense.Installments[2].PaidDate.Unix() - time.Now().Unix()

	expense.PayInstallment(2)

	assert.Less(t, TimeDifference, time.Minute)
	assert.Equal(t, expense.Installments[0].PaidDate, time.Time{})
	assert.Equal(t, expense.Installments[1].PaidDate, time.Time{})
	assert.Equal(t, expense.Installments[3].PaidDate, time.Time{})
}

func TestGivenAllInstallmentsToPay_WhenPayAllInstallments_ThenAllInstallmentsPaidSuceful(t *testing.T) {
	expense := NewExpense(data.Title, data.Description, data.Categories, data.Ammount, time.Now(), data.Paid, 12)

	expense.PayExpense()

	assert.NotEqual(t, expense.Installments[0].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[1].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[2].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[3].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[4].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[5].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[6].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[7].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[8].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[9].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[10].PaidDate, time.Time{})
	assert.NotEqual(t, expense.Installments[11].PaidDate, time.Time{})
	assert.True(t, expense.Paid)
}

func TestGivenPaidInstallment_WhenPayInstallment_ThenReturnError(t *testing.T) {
	expense := NewExpense(data.Title, data.Description, data.Categories, data.Ammount, time.Now(), data.Paid, 4)

	expense.PayInstallment(2)
	err := expense.PayInstallment(2)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "installment already paid")
}

func TestGivenPaidExpense_WhenPayExpense_ThenReturnError(t *testing.T) {
	data.Paid = true

	err := data.PayExpense()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "expense already paid")
}
