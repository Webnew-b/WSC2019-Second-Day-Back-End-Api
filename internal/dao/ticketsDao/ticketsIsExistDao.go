package ticketsDao

import (
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools/logUtil"
)

func TicketsIsExist(id int64, eventId int64) error {
	var ticket model.EventTickets

	data := database.GetDatabase().
		Where(&model.EventTickets{
			ID:      id,
			EventId: eventId,
		}).
		First(&ticket)

	if data.Error != nil {
		logUtil.Log.Println(data.Error.Error())
		return throwError()
	}

	ticketStatus := processEventTicket(ticket)
	if ticketStatus.Available {
		return nil
	}
	return throwError()
}
