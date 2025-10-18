package sqlite

import (
	"testing"
)

func TestHistoryModel(t *testing.T) {
	SetSqlite()

	// 同步表结构
	history := new(History)
	history.Sync()

	// 插入测试数据
	history.Title = "测试标题"
	err := history.Insert()
	if err != nil {
		t.Errorf("插入数据失败: %v", err)
	}
}
func TestExist(t *testing.T) {
	SetSqlite()
	history := new(History)
	history.Sync()

	// 插入测试数据
	history.Title = "测试标题"
	if has, err := history.ExistsByTitle(); has {
		t.Logf("数据%+v已存在", history)
	} else if err != nil {
		t.Errorf("查询数据失败: %v", err)
	} else {
		history.Insert()
	}
}
