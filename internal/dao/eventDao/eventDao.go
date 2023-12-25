package eventDao

import (
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
)

func GetAllEvent(events *[]model.Events) error {
	data := global.DB.Find(events)
	if data.Error != nil {
		return data.Error
	}
	return nil
}
