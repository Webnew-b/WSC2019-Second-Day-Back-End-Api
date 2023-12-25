package organizerDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
)

func GetOrganizerInfoById(id int64) (*api.ApiOrganizer, error) {
	var organizer api.ApiOrganizer
	data := global.DB.Model(&model.Organizers{}).Find(&organizer, id)
	if data.Error != nil {
		return nil, data.Error
	}
	return &organizer, nil
}
