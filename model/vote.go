package model

import "time"

type Vote struct {
	ID        uint
	Name      string `form:"name"`
	EventID   uint
	Scores    []Score
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (vote Vote) Create() (Vote, error) {
	err := db.Create(&vote).Error
	return vote, err
}
