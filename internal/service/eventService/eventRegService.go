package eventService

import (
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/dao/registrationsDao"
	"wscmakebygo.com/internal/dao/sessionsDao"
	"wscmakebygo.com/internal/dao/sessionsRegDao"
	"wscmakebygo.com/internal/dao/ticketsDao"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/internal/params/sessionParams"
)

const (
	success = "Registration successful"
)

func RegEvent(param *api.EventRegParams) (*api.EventRegRes, error) {
	var (
		event *api.EventDetailData
		err   error
	)
	event, err = fetchEvent(&eventParams.EventFetchRequest{
		OrgSlug: param.OrgSlug,
		EvSlug:  param.EvSlug,
	})
	if err != nil {
		return nil, err
	}

	err = ticketsDao.TicketsIsExist(param.TicketID, event.ID)
	if err != nil {
		return nil, err
	}

	err = sessionIsExist(param.SessionIds, event.ID)
	if err != nil {
		return nil, err
	}

	err = addReg(param)
	if err != nil {
		return nil, err
	}

	return buildSuccessMsg(), nil
}

func buildSuccessMsg() *api.EventRegRes {
	res := new(api.EventRegRes)
	res.Message = success
	return res
}

func sessionIsExist(sessionIds []int64, eventId int64) error {
	for _, id := range sessionIds {
		err := sessionsDao.SessionValid(id)
		if err != nil {
			return err
		}
		err = sessionsDao.IsSessionLinkedToEvent(id, eventId)
		if err != nil {
			return err
		}
	}
	return nil
}

func addReg(param *api.EventRegParams) error {
	var (
		sessionParam *sessionParams.SessionsRegCreate
		regId        int64
		err          error
	)
	tx := database.GetDatabase().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	regId, err = registrationsDao.AddRegistration(tx, param.AttendeeId, param.TicketID)
	if err != nil {
		tx.Rollback()
		return err
	}

	sessionParam = &sessionParams.SessionsRegCreate{
		Transaction:    tx,
		RegistrationId: regId,
	}

	err = regSessions(param.SessionIds, sessionParam)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func regSessions(sessions []int64, sessionParam *sessionParams.SessionsRegCreate) error {
	for _, sessionId := range sessions {
		sessionParam.SessionId = sessionId
		err := sessionsRegDao.AddSessionsReg(sessionParam)
		if err != nil {
			return err
		}
	}
	return nil
}
