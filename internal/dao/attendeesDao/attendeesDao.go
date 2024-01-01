package attendeesDao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/global/database"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/tools/logUtil"
	"wscmakebygo.com/tools/redisUtil"
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

func FetchAttendeeIdByCache(token string) (int64, error) {
	key := fmt.Sprintf("%s%s", constant.ATTENDEE_LOGIN_PREFIX, token)
	data, err := redisUtil.GetData(key)
	if err != nil {
		return 0, &attendeesError.NotLogin{}
	}
	id, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		logUtil.Log.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func checkedError(err error, msg string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logUtil.Log.Println(msg, "Attendees is not found")
		return &attendeesError.NotFound{}
	case err != nil:
		return err
	}
	return nil
}
