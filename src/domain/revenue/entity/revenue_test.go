package revenue

import (
	"strings"
	"testing"
	"time"

	"github.com/IsaqueAmorim/money-wise/src/domain/shared/category"
	assert "github.com/stretchr/testify/assert"
)

var rev = &revenue{
	Title:       "Title Test",
	Description: "Description Test",
	Ammount:     100.00,
	Categories:  []category.Category{},
	Date:        time.Now(),
}

func TestGivenValidParameters_WhenCreatingRevenue_ThenCreateExpenseWithParametersValue(t *testing.T) {

	dateBeforeCreate := time.Now()
	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Date(
			2003,
			time.August,
			13, 0, 0, 0, 0, time.Now().UTC().Location()))
	dateAfterCreate := time.Now()

	assert.NotEmpty(t, rev.ID.String())
	assert.Equal(t, rev.Title, "Title Test")
	assert.Equal(t, rev.Description, "Description Test")
	assert.Equal(t, rev.Ammount, 100.00)
	assert.NotEqual(t, rev.CreatedAt, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.LessOrEqual(t, rev.CreatedAt.Unix(), dateAfterCreate.Unix())
	assert.GreaterOrEqual(t, rev.CreatedAt.Unix(), dateBeforeCreate.Unix())
	assert.Equal(t, rev.Date, time.Date(2003, time.August, 13, 0, 0, 0, 0, time.Now().UTC().Location()))
}

func TestGivenValidTitle_WhenUpdateExpenseTitle_ThenChangeTitleSucess(t *testing.T) {

	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	err := rev.ChangeTitle("New Title")

	assert.Nil(t, err)
	assert.Equal(t, rev.Title, "New Title")
}

func TestGivenInvalidTitle_WhenUpdateExpenseTitle_ThenChangeTitleError(t *testing.T) {

	str := strings.Join(make([]string, 102), "A")
	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	errTitleLessThenTreeCharacters := rev.ChangeTitle("Ne")
	errTitleGreaterThenOneHundred := rev.ChangeTitle(str)

	assert.NotNil(t, errTitleLessThenTreeCharacters)
	assert.Equal(t, errTitleLessThenTreeCharacters.Error(), "title must be longer than 3 characters")
	assert.NotEqual(t, rev.Title, "Ne")

	assert.NotNil(t, errTitleGreaterThenOneHundred)
	assert.Equal(t, errTitleGreaterThenOneHundred.Error(), "title must be less than 100 characters")
	assert.NotEqual(t, rev.Title, str)

}

func TestGivenValidDescription_WhenUpdateExpenseDescription_ThenChangeDescriptionSucess(t *testing.T) {

	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	err := rev.ChangeDescription("New Description")

	assert.Nil(t, err)
	assert.Equal(t, rev.Description, "New Description")
	assert.NotEqual(t, rev.Description, "Description Test")
}

func TestGivenInvalidDescription_WhenUpdateExpenseDescription_ThenChangeDescriptionError(t *testing.T) {

	str := strings.Join(make([]string, 260), "A")
	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	err := rev.ChangeDescription(str)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "description must be less than 255 characters")
	assert.NotEqual(t, rev.Description, str)
	assert.Equal(t, rev.Description, "Description Test")
}

func TestGivenValidAmmout_WhenUpdatExpenseAmmout_ThenChangeAmount(t *testing.T) {

	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	err := rev.ChangeAmmount(200.00)

	assert.Nil(t, err)
	assert.Equal(t, rev.Ammount, 200.00)
	assert.NotEqual(t, rev.Ammount, 100.00)
}

func TestGivenInvalidAmount_WhenUpedadeAmount_ThenChangeAmountError(t *testing.T) {

	rev := NewRevenue(
		rev.Title,
		rev.Description,
		rev.Ammount,
		rev.Categories,
		time.Now())

	err := rev.ChangeAmmount(-200.00)
	err2 := rev.ChangeAmmount(0.00)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ammount must be greater than 0")
	assert.NotEqual(t, rev.Ammount, -200.00)
	assert.Equal(t, rev.Ammount, 100.00)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "ammount must be greater than 0")
	assert.NotEqual(t, rev.Ammount, 0.00)
}
