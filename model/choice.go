package model

import (
	"time"
)

type Choice struct {
	ID        uint
	Name      string
	EventID   uint
	Scores    []Score
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (choice Choice) Create() error {
	err := db.Create(&choice).Error
	return err
}
