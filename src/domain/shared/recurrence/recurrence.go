package recurrence

import (
	"time"

	"github.com/google/uuid"
)

type FrequencyEnum uint8

const (
	Daily FrequencyEnum = iota
	Weekly
	Fortnightly
	Monthly
	Quarterly
	SemiAnnual
	Annual
)

type Recurrence struct {
	ID          uuid.UUID
	FatherID    uuid.UUID
	Frequency   FrequencyEnum
	RepeatAt    time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	IsRecurring bool
}

func NewRecurrence(fatherID uuid.UUID, frequency FrequencyEnum, repeatAt, endDate time.Time, isRecurring bool) *Recurrence {
	return &Recurrence{
		ID:        uuid.New(),
		FatherID:  fatherID,
		Frequency: frequency,
		RepeatAt:  repeatAt,
		EndDate:   endDate,
		CreatedAt: time.Now(),
	}
}

func (r *Recurrence) ChangeFrequency(frequency FrequencyEnum) {
	r.Frequency = frequency
}

func (r *Recurrence) ChangeRepeatDate(repeatAt time.Time) {
	r.RepeatAt = repeatAt
}
