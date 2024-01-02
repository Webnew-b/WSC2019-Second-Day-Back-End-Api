package api

type FetchEventRegReq struct {
	Id int64 `validate:"required"`
}

type FetchEventRegRes struct {
	Registrations []Registrations `json:"registrations"`
}

type Registrations struct {
	Event      RegEvent `json:"event"`
	SessionIds []int64  `json:"session_ids"`
}

type RegEvent struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Date        *Date        `json:"date"`
	OrganizerId int64        `json:"-"`
	Organizer   RegOrganizer `json:"organizer" gorm:"-"`
}

type RegOrganizer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
