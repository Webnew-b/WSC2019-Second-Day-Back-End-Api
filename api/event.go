package api

import (
	"wscmakebygo.com/internal/model"
)

type EventsRes struct {
	Events []ApiEvent `json:"events"`
}

type ApiEvent struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Slug      string        `json:"slug"`
	Date      *model.Date   `json:"date"`
	Organizer *ApiOrganizer `json:"organizer"`
}
