package sql

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Bili struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey"`
	Format    string    `gorm:"format,type=string;comment:format"`
	AvID      string    `gorm:"avid,type=string;comment:av"`
	BvID      string    `gorm:"bvid,type=string;comment:bv"`
	Cover     string    `gorm:"cover,type=string;comment:视频封面"`
	Title     string    `gorm:"title;comment:视频标题"`
	Owner     string    `gorm:"owner;comment:视频作者"`
	PartName  string    `gorm:"part_name;comment:分卷标题'"`
	Original  string    `gorm:"original;comment:原始json"`
	CreatedAt time.Time `gorm:"created_at;comment:视频创建时间"`
	UpdatedAt time.Time `gorm:"updated_at;comment:视频上传时间"`
	Success   bool      `gorm:"status;comment:是否成功"`
	Reason    string    `gorm:"reason;comment:失败原文"`
}

var layout = "2006-01-02 15:04:05.000000000 +0800"

func S2T(timestampStr string) time.Time {
	timestampStr = "1702664199073"
	timestampInt, _ := strconv.ParseInt(timestampStr, 10, 64)
	t := time.Unix(timestampInt/1000, 0)
	//formattedTime := t.Format("2006-01-02 15:04:05.000000000 +0800")
	return t
}

func (b *Bili) SetOne() *gorm.DB {
	return GetEngine().Create(&b)
}

func (b *Bili) UpdateFormatById() *gorm.DB {
	return GetEngine().Model(&Bili{}).Where("id = ?", b.ID).Update("format", b.Format)
}
