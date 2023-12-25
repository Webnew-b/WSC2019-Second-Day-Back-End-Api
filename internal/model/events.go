package model

import (
	"fmt"
	"time"
)

type Date time.Time

func (t *Date) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}

type Events struct {
	ID          int64  `json:"id" gorm:"id"`
	OrganizerId int64  `json:"organizer_id" gorm:"organizer_id"`
	Name        string `json:"name" gorm:"name"`
	Slug        string `json:"slug" gorm:"slug"`
	Date        *Date  `json:"date" gorm:"date"`
}

// TableName 表名称
func (*Events) TableName() string {
	return "events"
}
