package attendeesDao

import (
	"errors"
	"gorm.io/gorm"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools"
)

func FetchAttendeesByCode(params api.LoginRequest) (*api.LoginRes, error) {
	var attendees api.LoginRes
	data := database.
		GetDatabase().
		Model(&model.Attendees{}).
		Where(&model.Attendees{
			Lastname:         params.Lastname,
			RegistrationCode: params.RegistrationCode,
		}, "lastname", "registration_code").
		First(&attendees)
	err := checkedError(data.Error, params.Lastname)
	if err != nil {
		return nil, err
	}
	return &attendees, nil
}

func checkedError(err error, msg string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		tools.Log.Println(msg, "Attendees is not found")
		return &attendeesError.NotFound{}
	case err != nil:
		return err
	}
	return nil
}
