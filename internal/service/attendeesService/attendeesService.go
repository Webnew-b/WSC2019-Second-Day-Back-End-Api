package attendeesService

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/dao/attendeesDao"
	"wscmakebygo.com/tools/logUtil"
	"wscmakebygo.com/tools/redisUtil"
)

func AttendeesLogin(params api.LoginRequest) (*api.LoginRes, error) {
	attendees, err := attendeesDao.FetchAttendeesByCode(params)
	if err != nil {
		return nil, err
	}
	token := createToken(attendees.Lastname)
	err = saveToken(token, attendees.Id)
	if err != nil {
		return nil, err
	}
	attendees.Token = token
	return attendees, nil
}

func AttendeesLogout(params api.LogoutRequest) (string, error) {
	err := removeToken(params.Token)
	if err != nil {
		return "", err
	}
	return "Logout success", nil
}

func createToken(lastname string) string {
	sliceName := []byte(lastname)
	token := md5.Sum(sliceName)
	hexToken := fmt.Sprintf("%x", token)
	hexToken = strings.ToUpper(hexToken)
	return hexToken
}

func saveToken(token string, id int64) error {
	key := fmt.Sprintf("%s%s", constant.ATTENDEE_LOGIN_PREFIX, token)
	err := redisUtil.SetData(key, strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}
	return nil
}

func removeToken(token string) error {
	key := fmt.Sprintf("%s%s", constant.ATTENDEE_LOGIN_PREFIX, token)
	err := redisUtil.RemoveKey(key)
	if err != nil {
		logUtil.Log.Println(err.Error())
		return &attendeesError.LoginKeyNotExist{}
	}
	return nil
}
