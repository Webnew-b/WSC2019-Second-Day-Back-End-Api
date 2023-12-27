package eventDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/internal/params/eventParams"
)

func GetAllEvent() (*[]model.Events, error) {
	var events []model.Events
	data := global.DB.Find(&events)
	if data.Error != nil {
		return nil, data.Error
	}
	return &events, nil
}

func GetEventDetail(param eventParams.EventFetchDao) (*api.EventDetailData, error) {
	var event api.EventDetailData
	data := global.DB.Model(&model.Events{}).Where(&model.Events{
		OrganizerId: param.OrgId,
		Slug:        param.EvSlug,
	}).First(&event)
	if data.Error != nil {
		//todo 要对找不到做一个处理。
		return nil, data.Error
	}
	return &event, nil
}
