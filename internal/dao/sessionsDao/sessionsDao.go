package sessionsDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
)

func FetchSessionsByRoomId(id int64) (*[]api.EventDetailSessions, error) {
	var rooms []api.EventDetailSessions
	data := global.DB.
		Model(&model.Sessions{}).
		Where(&model.Sessions{RoomId: id}).
		Find(&rooms)
	if data.Error != nil {
		return nil, data.Error
	}
	return &rooms, nil
}
