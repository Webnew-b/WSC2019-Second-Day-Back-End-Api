package eventDao

import (
	"errors"
	"gorm.io/gorm"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/tools"
)

func GetEventDetail(param eventParams.EventFetchDao) (*api.EventDetailData, error) {
	var event api.EventDetailData
	data := database.GetDatabase().Model(&model.Events{}).Where(&model.Events{
		OrganizerId: param.OrgId,
		Slug:        param.EvSlug,
	}).First(&event)
	err := checkedError(data.Error, param.EvSlug)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func FetchEventDetailByIds(ids []int64) (*[]api.RegEvent, error) {
	var event []api.RegEvent
	data := database.GetDatabase().Model(&model.Events{}).Where("id IN ?", ids).Find(&event)
	if data.Error != nil {
		return nil, checkedError(data.Error, "")
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
