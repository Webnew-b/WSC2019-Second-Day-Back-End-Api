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

func GetOrganizerIdBySlug(slug string) (int64, error) {
	var organizer api.ApiOrganizer
	data := global.DB.Model(&model.Organizers{}).Where(api.ApiOrganizer{Slug: slug}, "slug").First(&organizer)
	if data.Error != nil {
		// todo 要对找不到做一个处理
		return 0, data.Error
	}
	return organizer.ID, nil
}
