package ticketsDao

import (
	"fmt"
	"time"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/dao/registrationsDao"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools"
)

const (
	typeDate   = "date"
	typeAmount = "amount"
)

type specialValidity struct {
	Type       string `json:"type"`
	TicketDate string `json:"date"`
	Amount     int64  `json:"amount"`
}

func FetchTicketByEventId(id int64) (*[]api.EventDetailTickets, error) {
	var tickets []model.EventTickets
	data := global.DB.
		Where(&model.EventTickets{EventId: id}).Find(&tickets)
	if data.Error != nil {
		return nil, data.Error
	}
	res := buildEventDetailTickets(&tickets)
	return res, nil
}

func buildEventDetailTickets(tickets *[]model.EventTickets) *[]api.EventDetailTickets {
	list := make([]api.EventDetailTickets, len(*tickets))
	for index, ticket := range *tickets {
		var item specialValidity
		res := newEventDetailTicket(ticket)
		if ticket.SpecialValidity != "" {
			err := tools.JsonUnmarshal([]byte(ticket.SpecialValidity), &item)
			if err != nil {
				tools.Log.Println(err.Error())
				list[index] = *res
				continue
			}
		} else {
			res.Available = true
			list[index] = *res
			continue
		}
		res.Description = getDescription(item, ticket.ID)
		res.Available = getAvailable(item, ticket.ID)
		list[index] = *res
	}
	return &list
}

func getDescription(item specialValidity, ticketId int64) string {
	switch item.Type {
	case typeDate:
		date, err := parseDate(item.TicketDate)
		if err != nil {
			tools.Log.Println(err.Error())
			return fmt.Sprintf("Available until %s", item.TicketDate)
		}
		return fmt.Sprintf("Available until %s", formatDate(date))
	case typeAmount:
		if registrationsDao.CountTicketReg(ticketId) >= item.Amount {
			return "0 tickets available"
		}
		return fmt.Sprintf("%d tickets available", item.Amount)
	default:
		return "unknown"
	}
}

// todo 将这两个封装到tools里面去
func formatDate(date time.Time) string {
	newFormat := "January 2, 2006"
	formattedDate := date.Format(newFormat)
	return formattedDate
}

func parseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func getAvailable(item specialValidity, ticketId int64) bool {
	switch item.Type {
	case typeDate:
		date, err := parseDate(item.TicketDate)
		if err != nil {
			tools.Log.Println(err.Error())
			return false
		}
		now := time.Now()
		return date.After(now)
	case typeAmount:
		return registrationsDao.CountTicketReg(ticketId) >= item.Amount
	default:
		return false
	}
}

func newEventDetailTicket(ticket model.EventTickets) *api.EventDetailTickets {
	return &api.EventDetailTickets{
		ID:   ticket.ID,
		Name: ticket.Name,
		Cost: ticket.Cost,
	}
}
