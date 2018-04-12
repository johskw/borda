package service

import (
	"github.com/johskw/borda/model"
)

func CreateEventAndChoices(event model.Event) (model.Event, error) {
	event, err := event.Create()
	if err != nil {
		return event, err
	}
	for _, choiceName := range event.ChoiceNames {
		choice := model.Choice{
			Name:    choiceName,
			EventID: event.ID,
		}
		err = choice.Create()
		if err != nil {
			return event, err
		}
	}
	return event, nil
}
