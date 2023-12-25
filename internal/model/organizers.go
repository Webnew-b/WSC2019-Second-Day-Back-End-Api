package model

// Organizers undefined
type Organizers struct {
	ID           int64  `json:"id" gorm:"id"`
	Name         string `json:"name" gorm:"name"`
	Slug         string `json:"slug" gorm:"slug"`
	Email        string `json:"email" gorm:"email"`
	PasswordHash string `json:"password_hash" gorm:"password_hash"`
}

// TableName 表名称
func (*Organizers) TableName() string {
	return "organizers"
}
