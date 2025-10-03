package storage

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func SetMysql(user, password, host, port, dbName string) {
	// 第一次连接：不指定数据库，用于检查和创建数据库
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&loc=Asia%%2FShanghai", user, password, host, port)
	
	tempEngine, err := xorm.NewEngine("mysql", dsnWithoutDB)
	if err != nil {
		log.Fatalf("Fail to connect to MySQL server: %v", err)
	}
	defer tempEngine.Close()

	// 使用原始方法检查数据库是否存在
	var count int64
	sql := "SELECT COUNT(*) FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?"
	exists, err := tempEngine.SQL(sql, dbName).Get(&count)
	if err != nil {
		log.Fatalf("Fail to check database existence: %v", err)
	}
	
	if !exists || count == 0 {
		// 创建数据库
		_, err = tempEngine.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci'", dbName))
		if err != nil {
			log.Fatalf("Fail to create database %s: %v", dbName, err)
		}
		log.Printf("Database %s created successfully", dbName)
	}

	// 第二次连接：指定数据库进行正式连接
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai", user, password, host, port, dbName)
	
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