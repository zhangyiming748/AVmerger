package sqlite

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type History struct {
	Id       int64          `gorm:"primaryKey;autoIncrement;comment:主键"`
	Title    string         `gorm:"comment:标题"`
	CreateAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdateAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeleteAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}

func (h *History) Sync() {
	log.Printf("开始同步表结构\n")
	if err := GetSqlite().AutoMigrate(&History{}); err != nil {
		log.Fatalf("同步表结构History失败:%s", err.Error())
	}
	log.Printf("同步表结构完成\n")
}

func (h *History) Insert() error {
	db := GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Create(&h)
	return result.Error
}

func (h *History) Update() error {
	db := GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Save(&h)
	return result.Error
}

func (h *History) Delete() error {
	db := GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Delete(&h)
	return result.Error
}

func (h *History) GetById(id int64) error {
	db := GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.First(&h, id)
	return result.Error
}

func (h *History) GetAll() ([]History, error) {
	db := GetSqlite()
	if db == nil {
		return nil, errors.New("数据库连接未初始化")
	}
	var histories []History
	result := db.Find(&histories)
	return histories, result.Error
}

func (h *History) GetByTitle(title string) error {
	db := GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Where("title = ?", title).First(&h)
	return result.Error
}

func (h *History) ExistsByTitle() (bool, error) {
	db := GetSqlite()
	if db == nil {
		return false, errors.New("数据库连接未初始化")
	}
	var count int64
	result := db.Model(&History{}).Where("title = ?", h.Title).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}