package model

import "time"

type Events struct {
	ID          int64     `json:"id" gorm:"id"`
	OrganizerId int64     `json:"organizer_id" gorm:"organizer_id"`
	Name        string    `json:"name" gorm:"name"`
	Slug        string    `json:"slug" gorm:"slug"`
	Date        time.Time `json:"date" gorm:"date"`
}

// TableName 表名称
func (*Events) TableName() string {
	return "events"
}
