package model

import (
	"time"
)

type Vote struct {
	ID        uint
	Name      string   `form:"name"`
	ScoreStr  []string `form:"scores[]" gorm:"-"`
	Scores    []Score
	EventID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (vote Vote) StrEventID() string {
// 	return strconv.Itoa(int(vote.EventID))
// }

func (vote Vote) Create() (Vote, error) {
	err := db.Create(&vote).Error
	return vote, err
}
