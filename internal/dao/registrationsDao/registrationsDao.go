package registrationsDao

import (
	"gorm.io/gorm"
	"time"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/registrationsError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools/logUtil"
)

func CountTicketReg(id int64) int64 {
	var regs model.Registrations
	var count int64
	data := database.GetDatabase().
		Model(&regs).
		Where(&model.Registrations{TicketId: id}, "ticket_id").
		Count(&count)
	if data.Error != nil {
		logUtil.Log.Println(data.Error.Error())
		return 0
	}
	return count
}

func AddRegistration(tx *gorm.DB, attendeesId int64, ticketId int64) (int64, error) {
	var regs model.Registrations

	count, err := findRegByAttIdAndTicketId(attendeesId, ticketId)
	if err != nil {
		return 0, err
	} else if count > 0 {
		return 0, &registrationsError.AlreadyRegistrar{}
	}

	regs = model.Registrations{
		TicketId:         ticketId,
		AttendeeId:       attendeesId,
		RegistrationTime: time.Now(),
	}

	data := tx.Create(&regs)
	if data.Error != nil {
		logUtil.Log.Println(data.Error.Error())
		return -1, data.Error
	}

	return regs.ID, nil
}

func findRegByAttIdAndTicketId(attendeesId int64, ticketId int64) (int64, error) {
	var regs model.Registrations
	var count int64

	data := database.GetDatabase().
		Model(&regs).
		Where(&model.Registrations{
			AttendeeId: attendeesId,
			TicketId:   ticketId,
		}, "attendee_id", "ticket_id").
		Count(&count)

	if data.Error != nil {
		logUtil.Log.Println(data.Error.Error())
		return 0, data.Error
	}
	return count, nil
}
