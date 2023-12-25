package start

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wscmakebygo.com/global"
	"wscmakebygo.com/tools"
)

func crateDbAddr() string {
	Addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s&parseTime=true",
		global.Config.Db.User,
		global.Config.Db.Password,
		global.Config.Db.Host,
		global.Config.Db.Port,
		global.Config.Db.DbName,
		global.Config.Db.Charset,
		global.Config.Db.Loc,
	)
	return Addr
}

func crateDbConnect() {
	logStr := fmt.Sprintf("%s:%d", global.Config.Db.Host, global.Config.Db.Port)
	addr := crateDbAddr()
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	tools.Log.Println("created Db connection:" + logStr)
	if err != nil {
		panic(err)
	}
	global.DB = db
}
