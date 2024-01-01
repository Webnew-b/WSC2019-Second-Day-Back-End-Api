package eventService

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/internal/dao/attendeesDao"
	"wscmakebygo.com/internal/dao/eventDao"
	"wscmakebygo.com/internal/dao/organizerDao"
	"wscmakebygo.com/internal/dao/registrationsDao"
	"wscmakebygo.com/internal/dao/sessionsDao"
	"wscmakebygo.com/internal/dao/ticketsDao"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools/logUtil"
	"wscmakebygo.com/tools/uniqueUtil"
)

func FetchEventRegDetail(param *api.FetchEventRegReq) (*api.FetchEventRegRes, error) {
	var (
		event     *[]api.RegEvent
		regEvent  *[]api.RegEvent
		regs      *[]model.Registrations
		resReg    *[]api.Registrations
		attendId  int64
		ticketIds []int64
		eventIds  *[]int64
		err       error
	)

	logUtil.Log.Println(1)
	attendId, err = attendeesDao.FetchAttendeeIdByCache(param.Token)
	if err != nil {
		return nil, err
	}

	logUtil.Log.Println(2)
	regs, err = registrationsDao.FetchRegsByAttendeeId(attendId)
	if err != nil {
		return nil, err
	}

	ticketIds = fetchTicketId(*regs)

	logUtil.Log.Println(3)
	eventIds, err = fetchEventId(ticketIds)
	if err != nil {
		return nil, err
	}

	logUtil.Log.Println(4)
	event, err = eventDao.FetchEventDetailByIds(*eventIds)
	if err != nil {
		return nil, err
	}

	logUtil.Log.Println(5)
	regEvent, err = buildRegEvent(*event)
	if err != nil {
		return nil, err
	}

	logUtil.Log.Println(6)
	resReg, err = buildReg(*regEvent, attendId)

	return &api.FetchEventRegRes{Registrations: *resReg}, nil
}

func fetchTicketId(regs []model.Registrations) []int64 {
	var ticketIds []int64
	for _, reg := range regs {
		ticketIds = append(ticketIds, reg.TicketId)
	}
	return ticketIds
}

func fetchEventId(ticketIds []int64) (*[]int64, error) {
	uniqueTicketIds := uniqueUtil.UniqueInt64Slice(ticketIds)
	var eventIds []int64
	for _, id := range uniqueTicketIds {
		eventId, err := ticketsDao.FetchEventIdByTicketId(id)
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, eventId)
	}
	uniqueEventIds := uniqueUtil.UniqueInt64Slice(eventIds)
	return &uniqueEventIds, nil
}

func buildRegEvent(events []api.RegEvent) (*[]api.RegEvent, error) {
	for index := range events {
		org, err := organizerDao.GetRegOrganizerInfoById(events[index].OrganizerId)
		if err != nil {
			return nil, err
		}
		events[index].Organizer = *org
	}
	return &events, nil
}

func buildReg(events []api.RegEvent, attendeeId int64) (*[]api.Registrations, error) {
	var reg = make([]api.Registrations, len(events))
	for i, event := range events {
		sessionIds, err := sessionsDao.GetSessionIdsByEventIdAndAttendeeId(event.ID, attendeeId)
		if err != nil {
			return nil, err
		}
		reg[i] = api.Registrations{
			Event:      event,
			SessionIds: sessionIds,
		}
	}
	return &reg, nil
}
