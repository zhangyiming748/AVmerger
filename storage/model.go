package storage

import (
	"log"
	"time"
)

type History struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Title     string    `xorm:"comment('视频标题') VARCHAR(512)"`
	Keyword   string    `xorm:"comment('AVID or BVID') VARCHAR(512)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (h *History) Sync() {
	log.Printf("开始同步表结构")
	if err := GetMysql().Sync2(History{}); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}

func (h *History) InsertOne() (int64, error) {
	return GetMysql().InsertOne(h)
}
func (h *History) FindByTitle() (bool, error) {
	return GetMysql().Where("title = ?", h.Title).Get(h)
}
