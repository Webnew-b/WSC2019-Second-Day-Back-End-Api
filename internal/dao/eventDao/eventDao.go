package eventDao

import (
	"errors"
	"gorm.io/gorm"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/tools"
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
	err := checkedError(data.Error, param.EvSlug)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func checkedError(err error, msg string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		tools.Log.Println(msg, "event is not found")
		return &eventError.EventSlugNotFoundError{}
	case err != nil:
		return err
	}
	return nil
}
