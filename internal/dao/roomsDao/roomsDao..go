package roomsDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
)

func FetchRoomsByChannelId(id int64) (*[]api.EventDetailRooms, error) {
	var rooms []api.EventDetailRooms
	data := global.DB.
		Model(&model.Rooms{}).
		Select("id", "name").
		Where(&model.Rooms{ChannelId: id}).
		Find(&rooms)
	if data.Error != nil {
		return nil, data.Error
	}
	return &rooms, nil
}
