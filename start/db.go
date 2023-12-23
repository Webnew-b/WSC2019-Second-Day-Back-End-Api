package start

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"wscmakebygo.com/global"
)

func crateDbAddr() string {
	Addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s",
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
	log.Println("crate Db connection:" + logStr)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.DB = db
}
