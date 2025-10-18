package sqlite

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite() {
	err := os.MkdirAll("./data", os.ModePerm)
	if err != nil {
		log.Fatalf("创建本地sqlite数据库目录失败:%s", err.Error())
	}
	
	db, err := gorm.Open(sqlite.Open("./data/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("打开本地sqlite数据库失败:%s", err.Error())
	}
	
	gormDB = db
	log.Println("本地sqlite数据库初始化完成")
}

func GetSqlite() *gorm.DB {
	return gormDB
}