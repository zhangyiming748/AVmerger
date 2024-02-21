package sql

import (
	"gorm.io/gorm"
)

type Danmaku struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	Format  string `gorm:"format,type=string;comment:format"`
	AvID    string `gorm:"avid,type=string;comment:av"`
	BvID    string `gorm:"bvid,type=string;comment:bv"`
	Title   string `gorm:"title;comment:视频标题"`
	Content string `gorm:"content;comment:弹幕内容"`
}

func (d *Danmaku) SetOne() *gorm.DB {
	return GetEngine().Create(&d)
}
func (d *Danmaku) SetMany(list *[]Danmaku) *gorm.DB {
	return GetEngine().CreateInBatches(&list, 100000)
}
