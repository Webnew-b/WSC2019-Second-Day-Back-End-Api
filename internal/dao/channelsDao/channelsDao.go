package channelsDao

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/model"
)

func FetchChannelByEventId(id int64) (*[]api.EventDetailChannels, error) {
	var channels []api.EventDetailChannels
	data := database.GetDatabase().
		Model(&model.Channels{}).
		Select("id", "name").
		Where(&model.Channels{EventId: id}).Find(&channels)
	if data.Error != nil {
		return nil, data.Error
	}
	return &channels, nil
}
