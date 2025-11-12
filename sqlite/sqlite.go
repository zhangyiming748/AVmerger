package sqlite

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite() {
	home, _ := os.UserHomeDir()
	dbLocal := filepath.Join(home, "sqlite.db")
	db, err := gorm.Open(sqlite.Open(dbLocal), &gorm.Config{})
	if err != nil {
		log.Fatalf("打开本地sqlite数据库失败:%s", err.Error())
	}

	gormDB = db
	log.Printf("本地sqlite数据库初始化完成!位置在%s\n", dbLocal)
}

func GetSqlite() *gorm.DB {
	return gormDB
}
