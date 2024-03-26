package expense

import (
	"github.com/IsaqueAmorim/money-wise/domain/shared/category"
	"github.com/google/uuid"
	"time"
)

type expense struct {
	ID          uuid.UUID
	Title       string
	Description string
	Categories  []category.Category
	Ammount     float64
	CreatedAt   time.Time
	Date        time.Time
}

func NewExpense(title string, description string, categories []category.Category, ammount float64, date time.Time) *expense {
	return &expense{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Categories:  categories,
		Ammount:     ammount,
		CreatedAt:   time.Now(),
		Date:        date,
	}
}
