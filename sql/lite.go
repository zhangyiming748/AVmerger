package sql

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func SetEngine() {
	db, _ = gorm.Open(sqlite.Open("merge.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	// 迁移 schema
	err := db.AutoMigrate(Bili{})
	if err != nil {
		panic("创建数据库错误")
		return
	}
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	//db.Delete(&product, 1)
	fmt.Println(db)
}
func GetEngine() *gorm.DB {
	return db
}
