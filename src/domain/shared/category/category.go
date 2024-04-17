package category

import "github.com/google/uuid"

const (
	EXPENSE = "EXPENSE"
	REVENUE = "REVENUE"
)

type Category struct {
	ID   uuid.UUID
	Name string
	Type string
}

func NewExpenseCategory(name string) *Category {
	return &Category{
		ID:   uuid.New(),
		Name: name,
		Type: EXPENSE,
	}
}

func NewRevenueCategory(name string) *Category {
	return &Category{
		ID:   uuid.New(),
		Name: name,
		Type: REVENUE,
	}
}
