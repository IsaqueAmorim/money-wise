package revenue

import (
	"github.com/IsaqueAmorim/money-wise/domain/shared/category"
	"github.com/google/uuid"
	"time"
)

type Revenue struct {
	ID          uuid.UUID
	Title       string
	Description string
	Ammount     float64
	Categories  []category.Category
	CreatedAt   time.Time
	Date        time.Time
}

func NewRevenue(ID string, title string, description string, ammount float64, categories []category.Category, date time.Time) *Revenue {
	return &Revenue{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Ammount:     ammount,
		Categories:  categories,
		CreatedAt:   time.Now(),
		Date:        date,
	}
}
