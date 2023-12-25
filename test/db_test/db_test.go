package db_test

import (
	"log"
	"testing"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/start"
	"wscmakebygo.com/test"
	"wscmakebygo.com/tools"
)

var (
	event  model.Events
	events []model.Events
)

func TestDb(t *testing.T) {
	test.ChangeDir()
	start.StartDbConnect()
	data := global.DB.Select("slug", "name").Find(&events)
	if data.Error != nil {
		tools.Log.Println("123456")
		tools.Log.Fatal(data.Error, "123456")
	}
	str, err := tools.JsonMarshal(events)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(str))
}

func TestGetDirPath(t *testing.T) {

}
