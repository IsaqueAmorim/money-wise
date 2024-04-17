package revenue

import (
	"time"

	"github.com/IsaqueAmorim/money-wise/src/domain/shared/category"
	"github.com/google/uuid"
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
