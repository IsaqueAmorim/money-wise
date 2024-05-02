package revenue

import (
	"time"

	"github.com/IsaqueAmorim/money-wise/src/domain/shared/category"
	"github.com/IsaqueAmorim/money-wise/src/domain/shared/recurrence"
	"github.com/google/uuid"
)

type Revenue struct {
	ID          uuid.UUID
	Title       string
	Description string
	Ammount     float64
	Categories  []category.Category
	Recurrence  recurrence.Recurrence
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

func (r *Revenue) ChangeTitle(title string) {
	r.Title = title
}

func (r *Revenue) ChangeDescription(description string) {
	r.Description = description
}

func (r *Revenue) ChangeAmmount(ammount float64) {
	r.Ammount = ammount
}

func (r *Revenue) ChangeCategories(categories []category.Category) {
	r.Categories = categories
}
