package eventDao

import (
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/model"
)

func GetAllEvent() (*[]model.Events, error) {
	var events []model.Events
	data := database.GetDatabase().Find(&events)
	if data.Error != nil {
		return nil, data.Error
	}
	return &events, nil
}
