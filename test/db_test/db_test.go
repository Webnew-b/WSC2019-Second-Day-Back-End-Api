package db_test

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"testing"
	"wscmakebygo.com/api"
	"wscmakebygo.com/global"
	"wscmakebygo.com/internal/model"
	"wscmakebygo.com/start"
	"wscmakebygo.com/test"
	"wscmakebygo.com/tools"
)

var (
	event  model.Events
	events []model.Events
	ticket api.EventDetailTickets
)

func TestDb(t *testing.T) {
	/*test.ChangeDir()
	start.StartDbConnect()
	data := global.DB.Model(&model.EventTickets{}).Find(&ticket)
	if data.Error != nil {
		tools.Log.Fatal(data.Error.Error())
	}
	tools.Log.Println(ticket.Cost)*/
	str, err := tools.JsonMarshal(ticket)
	if err != nil {
		log.Fatal(err)
	}
	tools.Log.Println(string(str))
}

func TestGetDirPath(t *testing.T) {
	test.ChangeDir()
	start.StartDbConnect()
	data := global.DB.First(&event, 6)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		// 如果没有找到记录，这里可以处理错误，比如抛出一个自定义错误或进行日志记录
		log.Fatalf("Error: Record not found for user with ID %d", 6)
	}
	if data.Error != nil {
		tools.Log.Println("123456")
		tools.Log.Fatal(data.Error, "123456")
	}
	str, err := tools.JsonMarshal(event)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(str))
}

func TestLogPrint(t *testing.T) {
	log.Println("aaa", "bbb")
}
