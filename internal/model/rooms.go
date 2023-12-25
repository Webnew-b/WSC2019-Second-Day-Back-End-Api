package model

type Rooms struct {
	ID        int64  `json:"id" gorm:"id"`
	ChannelId int64  `json:"channel_id" gorm:"channel_id"`
	Name      string `json:"name" gorm:"name"`
	Capacity  int64  `json:"capacity" gorm:"capacity"`
}

// TableName 表名称
func (*Rooms) TableName() string {
	return "rooms"
}
