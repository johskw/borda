package model

import "time"

type Score struct {
	ID        uint
	Score     uint `json:"score"`
	ChoiceID  uint `json:"choice_id"`
	VoteID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (score Score) Create() error {
	err := db.Create(&score).Error
	return err
}
