package sessionsDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/model"
)

func FetchSessionsByRoomId(id int64) (*[]api.EventDetailSessions, error) {
	var rooms []api.EventDetailSessions
	data := database.GetDatabase().
		Model(&model.Sessions{}).
		Where(&model.Sessions{RoomId: id}).
		Find(&rooms)
	if data.Error != nil {
		return nil, data.Error
	}
	return &rooms, nil
}
