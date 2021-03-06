package model

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Event struct {
	ID          uint
	Theme       string `form:"theme"`
	Detail      string `form:"detail"`
	Password    string `form:"password"`
	Finished    bool
	ChoiceNames []string `form:"choices[]" gorm:"-"`
	Choices     []Choice
	Votes       []Vote
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
	err := db.First(&event, id).Related(&event.Choices).Related(&event.Votes).Error
	return event, err
}

func (event Event) Update() error {
	err := db.Save(&event).Error
	return err
}

func (event Event) Finish() error {
	event.Finished = true
	err := db.Save(&event).Error
	return err
}

func (event Event) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(event.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
