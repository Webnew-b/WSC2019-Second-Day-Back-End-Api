package model

type Channels struct {
	ID      int64  `json:"id" gorm:"id"`
	EventId int64  `json:"event_id" gorm:"event_id"`
	Name    string `json:"name" gorm:"name"`
}

// TableName 表名称
func (*Channels) TableName() string {
	return "channels"
}
