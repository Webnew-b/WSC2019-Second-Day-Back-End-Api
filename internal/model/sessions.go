package model

import "time"

type Sessions struct {
	ID          int64     `json:"id" gorm:"id"`
	RoomId      int64     `json:"room_id" gorm:"room_id"`
	Title       string    `json:"title" gorm:"title"`
	Description string    `json:"description" gorm:"description"`
	Speaker     string    `json:"speaker" gorm:"speaker"`
	Start       time.Time `json:"start" gorm:"start"`
	End         time.Time `json:"end" gorm:"end"`
	Type        string    `json:"type" gorm:"type"`
	Cost        float64   `json:"cost" gorm:"cost"`
}

// TableName 表名称
func (*Sessions) TableName() string {
	return "sessions"
}
