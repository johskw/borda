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

type Scores []Score

func (score Score) Create() error {
	err := db.Create(&score).Error
	return err
}

func GetScoreByChoiceID(choiceID uint) (Scores, error) {
	var scores Scores
	err := db.Where("choice_id = ?", choiceID).Find(&scores).Error
	return scores, err
}

func (scores Scores) TotalScore() uint {
	var totalScore uint
	for _, score := range scores {
		totalScore += score.Score
	}
	return totalScore
}
