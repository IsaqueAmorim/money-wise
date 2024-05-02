package category

import "github.com/google/uuid"

type TransactionTypeEnum uint8

const (
	EXPENSE TransactionTypeEnum = iota
	REVENUE
)

type Category struct {
	ID   uuid.UUID
	Name string
	Type TransactionTypeEnum
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
