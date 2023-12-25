package model

// EventTickets undefined
type EventTickets struct {
	ID              int64   `json:"id" gorm:"id"`
	EventId         int64   `json:"event_id" gorm:"event_id"`
	Name            string  `json:"name" gorm:"name"`
	Cost            float64 `json:"cost" gorm:"cost"`
	SpecialValidity string  `json:"special_validity" gorm:"special_validity"`
}

// TableName 表名称
func (*EventTickets) TableName() string {
	return "event_tickets"
}
