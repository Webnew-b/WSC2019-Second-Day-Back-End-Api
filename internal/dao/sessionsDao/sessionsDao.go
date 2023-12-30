package sessionsDao

import (
	"fmt"
	"time"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/ticketsError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools"
)

func FetchSessionsByRoomId(id int64) (*[]api.EventDetailSessions, error) {
	var rooms []api.EventDetailSessions
	data := database.GetDatabase().
		Model(&model.Sessions{}).
		Where(&model.Sessions{RoomId: id}).
		Find(&rooms)
	if data.Error != nil {
		return nil, data.Error
	}
	return &rooms, nil
}

func SessionValid(id int64) error {
	var session model.Sessions
	data := database.GetDatabase().First(&session, id)
	if data.Error != nil {
		tools.Log.Println(data.Error.Error(), fmt.Sprintf("sessionId:%d", id))
		return throwError()
	}
	if isOutTime(session.End) {
		tools.Log.Println(fmt.Sprintf("sessionId:%d", id), "session is out Time")
		return throwError()
	}
	return nil
}

func isOutTime(date time.Time) bool {
	now := time.Now()
	return date.Before(now)
}

func throwError() error {
	return &ticketsError.NotAvailable{}
}
