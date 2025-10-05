package storage

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func SetMysql(user, password, host, port string, dbName string) {
	// 直接连接到指定的数据库merge
	dsnWithDB := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=Asia%%2FShanghai", user, password, host, port, dbName)

	var err error
	engine, err = xorm.NewEngine("mysql", dsnWithDB)
	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}

	if err := engine.Ping(); err != nil {
		log.Fatalf("连接数据库出错:%v\n", err)
	}

	// 可选：显示 SQL 语句
	engine.ShowSQL(true)

	// 可选：设置连接池
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)

	log.Println("Successfully connected to database!")
}

func GetMysql() *xorm.Engine {
	return engine
}
