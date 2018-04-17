package service

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/johskw/borda/model"
)

func CreateEventAndChoices(event model.Event) (model.Event, error) {
	event.Finished = false
	hash, err := bcrypt.GenerateFromPassword([]byte(event.Password), bcrypt.DefaultCost)
	if err != nil {
		return event, err
	}
	event.Password = string(hash[:])
	event, err = event.Create()
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

func GetEventFromStrID(strID string) (model.Event, error) {
	intID, _ := strconv.Atoi(strID)
	event, err := model.GetEvent(uint(intID))
	return event, err
}

func FinishEvent(strID string) error {
	event, err := GetEventFromStrID(strID)
	if err != nil {
		return err
	}
	err = event.Finish()
	return err
}
