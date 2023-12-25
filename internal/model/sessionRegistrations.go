package model

type SessionRegistrations struct {
	ID             int64 `json:"id" gorm:"id"`
	RegistrationId int64 `json:"registration_id" gorm:"registration_id"`
	SessionId      int64 `json:"session_id" gorm:"session_id"`
}

// TableName 表名称
func (*SessionRegistrations) TableName() string {
	return "session_registrations"
}
