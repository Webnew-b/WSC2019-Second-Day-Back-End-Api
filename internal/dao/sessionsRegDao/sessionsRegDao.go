package sessionsRegDao

import (
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/internal/params/sessionParams"
)

func AddSessionsReg(sessionParam *sessionParams.SessionsRegCreate) error {
	sessionReg := model.SessionRegistrations{
		SessionId:      sessionParam.SessionId,
		RegistrationId: sessionParam.RegistrationId,
	}

	data := sessionParam.Transaction.Create(&sessionReg)

	if data.Error != nil {
		return data.Error
	}
	return nil
}
