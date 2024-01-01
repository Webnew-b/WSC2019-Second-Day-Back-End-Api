package model

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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

func (*Attendees) CheckAttendeeType(value interface{}) (Attendees, error) {
	attendee, ok := value.(Attendees)
	if !ok {
		return Attendees{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid attendee type")
	}
	return attendee, nil
}
