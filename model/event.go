package model

import (
	"strconv"
	"time"
)

type Event struct {
	ID          uint
	Title       string   `form:"title"`
	ChoiceNames []string `form:"choices[]" gorm:"-"`
	Choices     []Choice
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (event Event) StrID() string {
	return strconv.Itoa(int(event.ID))
}

func (event Event) Create() (Event, error) {
	err := db.Create(&event).Error
	return event, err
}
