package eventService

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/dao/channelsDao"
	"wscmakebygo.com/internal/dao/eventDao"
	"wscmakebygo.com/internal/dao/organizerDao"
	"wscmakebygo.com/internal/dao/roomsDao"
	"wscmakebygo.com/internal/dao/sessionsDao"
	"wscmakebygo.com/internal/dao/ticketsDao"
	"wscmakebygo.com/internal/params/eventParams"
)

func FetchEventDetail(param eventParams.EventFetchRequest) (*api.EventDetailRes, error) {
	var (
		event *api.EventDetailData
		res   *api.EventDetailRes

		err error
	)

	event, err = fetchEvent(&param)
	if err != nil {
		return nil, err
	}

	res, err = buildEventDetailRes(event)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func fetchEvent(param *eventParams.EventFetchRequest) (*api.EventDetailData, error) {
	var (
		orgId int64
		event *api.EventDetailData

		err error
	)
	orgId, err = organizerDao.GetOrganizerIdBySlug(param.OrgSlug)
	if err != nil {
		return nil, err
	}

	fetchDao := eventParams.EventFetchDao{
		EventFetchRequest: param,
		OrgId:             orgId,
	}

	event, err = eventDao.GetEventDetail(fetchDao)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func buildEventDetailRes(event *api.EventDetailData) (*api.EventDetailRes, error) {
	res := api.EventDetailRes{
		EventDetailData: event,
	}

	err := fetchChannels(&res)
	if err != nil {
		return nil, err
	}
	err = fetchRooms(&res)
	if err != nil {
		return nil, err
	}
	err = fetchTickets(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func fetchChannels(event *api.EventDetailRes) error {
	channels, err := channelsDao.FetchChannelByEventId(event.ID)
	if err != nil {
		return err
	}
	event.Channels = *channels
	return err
}

func fetchTickets(event *api.EventDetailRes) error {
	tickets, err := ticketsDao.FetchTicketByEventId(event.ID)
	if err != nil {
		return err
	}
	event.Tickets = *tickets
	return err
}

func fetchRooms(event *api.EventDetailRes) error {
	for index := range event.Channels {
		rooms, err := roomsDao.FetchRoomsByChannelId((event.Channels)[index].ID)
		if err != nil {
			return err
		}

		err = fetchSession(rooms)
		if err != nil {
			return err
		}
		(event.Channels)[index].Rooms = *rooms
	}
	return nil
}

func fetchSession(rooms *[]api.EventDetailRooms) error {
	for index := range *rooms {
		sessions, err := sessionsDao.FetchSessionsByRoomId((*rooms)[index].ID)
		if err != nil {
			return err
		}
		(*rooms)[index].Sessions = *sessions
	}
	return nil
}
