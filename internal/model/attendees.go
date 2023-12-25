package model

type Attendees struct {
	ID               int64  `json:"id" gorm:"id"`
	Firstname        string `json:"firstname" gorm:"firstname"`
	Lastname         string `json:"lastname" gorm:"lastname"`
	Username         string `json:"username" gorm:"username"`
	Email            string `json:"email" gorm:"email"`
	RegistrationCode string `json:"registration_code" gorm:"registration_code"`
	LoginToken       string `json:"login_token" gorm:"login_token"`
}

// TableName 表名称
func (*Attendees) TableName() string {
	return "attendees"
}
