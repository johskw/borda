package service

import (
	"sort"

	"github.com/johskw/borda/model"
)

type ResultItem struct {
	ChoiceName string
	TotalScore uint
}

func GetResult(event model.Event) ([]ResultItem, error) {
	var result []ResultItem
	for _, choice := range event.Choices {
		scores, err := model.GetScoreByChoiceID(choice.ID)
		if err != nil {
			return result, err
		}
		resultItem := ResultItem{
			ChoiceName: choice.Name,
			TotalScore: scores.TotalScore(),
		}
		result = append(result, resultItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TotalScore > result[j].TotalScore
	})
	return result, nil
}
