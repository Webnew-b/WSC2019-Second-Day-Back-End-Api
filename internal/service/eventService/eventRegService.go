package eventService

import (
	"fmt"
	"strconv"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/params/eventParams"
	"wscmakebygo.com/tools"
	"wscmakebygo.com/tools/redisUtil"
)

func RegEvent(param *api.EventRegParams) (*api.EventRegRes, error) {
	var (
		event    *api.EventDetailData
		res      *api.EventRegRes
		attendId int64

		err error
	)
	event, err = fetchEvent(&eventParams.EventFetchRequest{
		OrgSlug: param.OrgSlug,
		EvSlug:  param.EvSlug,
	})
	if err != nil {
		return nil, err
	}

	attendId, err = fetchAttendeeId(param.Token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func fetchAttendeeId(token string) (int64, error) {
	key := fmt.Sprintf("%s%s", constant.ATTENDEE_LOGIN_PREFIX, token)
	data, err := redisUtil.GetData(key)
	if err != nil {
		return 0, &attendeesError.NotLogin{}
	}
	id, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		tools.Log.Println(err.Error())
		return 0, err
	}
	return id, nil
}
