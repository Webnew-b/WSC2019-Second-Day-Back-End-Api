package api

type EventDetailData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Date *Date  `json:"date"`
}

type EventDetailRes struct {
	*EventDetailData
	Channels []EventDetailChannels `json:"channels"`
	Tickets  []EventDetailTickets  `json:"tickets"`
}

type EventDetailSessions struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Speaker     string    `json:"speaker"`
	Start       *DateTime `json:"start"`
	End         *DateTime `json:"end"`
	Type        string    `json:"type"`
	Cost        float64   `json:"cost"`
}

type EventDetailRooms struct {
	ID       int64                 `json:"id"`
	Name     string                `json:"name"`
	Sessions []EventDetailSessions `json:"sessions" gorm:"-"`
}
type EventDetailChannels struct {
	ID    int64              `json:"id"`
	Name  string             `json:"name"`
	Rooms []EventDetailRooms `json:"rooms" gorm:"-"`
}

type EventDetailTickets struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	Available   bool    `json:"available"`
}
