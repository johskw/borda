package model

import (
	"strconv"
	"time"
)

type Event struct {
	ID          uint
	Theme       string   `form:"theme"`
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

func GetEvent(id uint) (Event, error) {
	var event Event
	err := db.First(&event, id).Related(&event.Choices).Error
	return event, err
}
