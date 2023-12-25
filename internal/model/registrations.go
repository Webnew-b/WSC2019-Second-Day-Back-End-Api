package model

import "time"

type Registrations struct {
	ID               int64     `json:"id" gorm:"id"`
	AttendeeId       int64     `json:"attendee_id" gorm:"attendee_id"`
	TicketId         int64     `json:"ticket_id" gorm:"ticket_id"`
	RegistrationTime time.Time `json:"registration_time" gorm:"registration_time"`
}

// TableName 表名称
func (*Registrations) TableName() string {
	return "registrations"
}
