package registrationsDao

import (
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools"
)

func CountTicketReg(id int64) int64 {
	var regs model.Registrations
	var count int64
	data := global.DB.
		Model(&regs).
		Where(&model.Registrations{TicketId: id}).
		Count(&count)
	if data.Error != nil {
		tools.Log.Println(data.Error.Error())
		return 0
	}
	return count
}
