package eventService

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/dao/eventDao"
	"wscmakebygo.com/internal/dao/organizerDao"
	"wscmakebygo.com/internal/model"
)

func GetAllEventAndOrganizer() (*[]api.ApiEvent, error) {
	var events []model.Events
	var apiEvents *[]api.ApiEvent
	err := eventDao.GetAllEvent(&events)
	if err != nil {
		return nil, err
	}
	apiEvents, err = buildApiEvent(&events)
	if err != nil {
		return nil, err
	}
	return apiEvents, nil
}

func buildApiEvent(events *[]model.Events) (*[]api.ApiEvent, error) {
	var apiEvents []api.ApiEvent
	for _, event := range *events {
		organizer, err := organizerDao.GetOrganizerInfoById(event.OrganizerId)
		if err != nil {
			return nil, err
		}
		apiEvent := api.ApiEvent{
			ID:        event.ID,
			Name:      event.Name,
			Slug:      event.Slug,
			Date:      event.Date,
			Organizer: organizer,
		}
		apiEvents = append(apiEvents, apiEvent)
	}
	return &apiEvents, nil
}
