package registrationsDao

import (
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/model"
)

func FetchRegsByAttendeeId(attendeeId int64) (*[]model.Registrations, error) {
	var regs []model.Registrations
	data := database.
		GetDatabase().
		Where(&model.Registrations{
			AttendeeId: attendeeId,
		}, "attendee_id").
		Find(&regs)
	if data.Error != nil {
		return &regs, data.Error
	}
	return &regs, nil
}
