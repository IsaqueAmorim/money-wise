package revenue

import (
	"errors"
	"time"

	"github.com/IsaqueAmorim/money-wise/src/domain/shared/category"
	"github.com/IsaqueAmorim/money-wise/src/domain/shared/recurrence"
	"github.com/google/uuid"
)

type revenue struct {
	ID          uuid.UUID
	Title       string
	Description string
	Ammount     float64
	Categories  []category.Category
	Recurrence  recurrence.Recurrence
	CreatedAt   time.Time
	Date        time.Time
}

func NewRevenue(title, description string, ammount float64, categories []category.Category, date time.Time) *revenue {
	return &revenue{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Ammount:     ammount,
		Categories:  categories,
		CreatedAt:   time.Now(),
		Date:        date,
	}
}

func (r *revenue) ChangeTitle(title string) error {
	if len(title) <= 3 {
		return errors.New("title must be longer than 3 characters")
	}
	if len(title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	r.Title = title
	return nil
}

func (r *revenue) ChangeDescription(description string) error {
	if len(description) > 255 {
		return errors.New("description must be less than 255 characters")
	}
	r.Description = description
	return nil
}

func (r *revenue) ChangeAmmount(ammount float64) error {
	if ammount <= 0 {
		return errors.New("ammount must be greater than 0")
	}
	r.Ammount = ammount
	return nil
}
