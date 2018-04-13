package service

import (
	"encoding/json"

	"github.com/johskw/borda/model"
)

func CreateVoteAndScores(vote model.Vote) (model.Vote, error) {
	vote, err := vote.Create()
	if err != nil {
		return vote, err
	}
	for _, scoreStr := range vote.ScoreStr {
		score := model.Score{
			VoteID: vote.ID,
		}
		err := json.Unmarshal([]byte(scoreStr), &score)
		if err != nil {
			return vote, err
		}
		err = score.Create()
		if err != nil {
			return vote, err
		}
	}
	return vote, nil
}
