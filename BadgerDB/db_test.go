package badgerdb

import (
	"os"
	"testing"
	"time"
)

// 测试数据库目录
const testDBDir = "./testdata/testdb"

// setupTestDB 创建测试数据库
func setupTestDB(t *testing.T) *DB {
	// 清理旧的测试数据
	os.RemoveAll(testDBDir)

	config := DefaultConfig(testDBDir)
	db, err := New(config)
	if err != nil {
		t.Fatalf("创建测试数据库失败: %v", err)
	}

	return db
}

// teardownTestDB 清理测试数据库
func teardownTestDB(t *testing.T, db *DB) {
	if err := db.Close(); err != nil {
		t.Fatalf("关闭数据库失败: %v", err)
	}
	os.RemoveAll(testDBDir)
}

// TestNew 测试数据库初始化
func TestNew(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	if db == nil {
		t.Fatal("数据库实例不能为空")
	}

	if db.IsClosed() {
		t.Fatal("新创建的数据库不应处于关闭状态")
	}
}

// TestSetAndGet 测试设置和获取键值对
func TestSetAndGet(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	type TestData struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	// 测试设置和获取
	data := TestData{Name: "test", Value: 123}
	err := db.Set("key1", data, 0)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	var retrievedData TestData
	err = db.Get("key1", &retrievedData)
	if err != nil {
		t.Fatalf("获取键值对失败: %v", err)
	}

	if retrievedData.Name != data.Name || retrievedData.Value != data.Value {
		t.Errorf("获取的数据与设置的不匹配: 期望 %+v, 得到 %+v", data, retrievedData)
	}
}

// TestGetNotFound 测试获取不存在的键
func TestGetNotFound(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	var data string
	err := db.Get("nonexistent", &data)
	if err == nil {
		t.Fatal("获取不存在的键应该返回错误")
	}
}

// TestDelete 测试删除键值对
func TestDelete(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 先设置一个键
	err := db.Set("key1", "value1", 0)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 删除该键
	err = db.Delete("key1")
	if err != nil {
		t.Fatalf("删除键值对失败: %v", err)
	}

	// 验证键已被删除
	var value string
	err = db.Get("key1", &value)
	if err == nil {
		t.Fatal("删除后的键不应该存在")
	}
}

// TestExists 测试检查键是否存在
func TestExists(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 设置一个键
	err := db.Set("key1", "value1", 0)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 检查键是否存在
	exists, err := db.Exists("key1")
	if err != nil {
		t.Fatalf("检查键存在性失败: %v", err)
	}
	if !exists {
		t.Error("键应该存在")
	}

	// 检查不存在的键
	exists, err = db.Exists("nonexistent")
	if err != nil {
		t.Fatalf("检查键存在性失败: %v", err)
	}
	if exists {
		t.Error("键不应该存在")
	}
}

// TestBatchSet 测试批量设置
func TestBatchSet(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 验证所有键都已设置
	for key, expectedValue := range items {
		var value string
		err := db.Get(key, &value)
		if err != nil {
			t.Fatalf("获取键 %s 失败: %v", key, err)
		}
		if value != expectedValue {
			t.Errorf("键 %s 的值不匹配: 期望 %s, 得到 %s", key, expectedValue, value)
		}
	}
}

// TestBatchDelete 测试批量删除
func TestBatchDelete(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 先设置一些键
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 批量删除
	keys := []string{"key1", "key2"}
	err = db.BatchDelete(keys)
	if err != nil {
		t.Fatalf("批量删除失败: %v", err)
	}

	// 验证删除的键不存在
	for _, key := range keys {
		var value string
		err := db.Get(key, &value)
		if err == nil {
			t.Errorf("键 %s 应该被删除", key)
		}
	}

	// 验证未删除的键仍然存在
	var value string
	err = db.Get("key3", &value)
	if err != nil {
		t.Errorf("键 key3 应该仍然存在: %v", err)
	}
}

// TestGetAllKeys 测试获取所有键
func TestGetAllKeys(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 设置一些键
	items := map[string]interface{}{
		"user:1":    "Alice",
		"user:2":    "Bob",
		"session:1": "active",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 获取所有键
	keys, err := db.GetAllKeys()
	if err != nil {
		t.Fatalf("获取所有键失败: %v", err)
	}

	if len(keys) != 3 {
		t.Errorf("期望 3 个键, 得到 %d 个", len(keys))
	}
}

// TestGetKeysWithPrefix 测试获取指定前缀的键
func TestGetKeysWithPrefix(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 设置一些键
	items := map[string]interface{}{
		"user:1":    "Alice",
		"user:2":    "Bob",
		"user:3":    "Charlie",
		"session:1": "active",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 获取 user: 前缀的键
	keys, err := db.GetKeysWithPrefix("user:")
	if err != nil {
		t.Fatalf("获取前缀键失败: %v", err)
	}

	if len(keys) != 3 {
		t.Errorf("期望 3 个 user: 前缀的键, 得到 %d 个", len(keys))
	}
}

// TestCount 测试统计键数量
func TestCount(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 初始计数应为 0
	count, err := db.Count()
	if err != nil {
		t.Fatalf("统计失败: %v", err)
	}
	if count != 0 {
		t.Errorf("期望 0 个键, 得到 %d 个", count)
	}

	// 添加一些键
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	err = db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 再次统计
	count, err = db.Count()
	if err != nil {
		t.Fatalf("统计失败: %v", err)
	}
	if count != 2 {
		t.Errorf("期望 2 个键, 得到 %d 个", count)
	}
}

// TestClear 测试清空数据库
func TestClear(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 添加一些键
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 清空数据库
	err = db.Clear()
	if err != nil {
		t.Fatalf("清空数据库失败: %v", err)
	}

	// 验证数据库已清空
	count, err := db.Count()
	if err != nil {
		t.Fatalf("统计失败: %v", err)
	}
	if count != 0 {
		t.Errorf("清空后期望 0 个键, 得到 %d 个", count)
	}
}

// TestTTL 测试过期时间
func TestTTL(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 设置一个带短过期时间的键
	err := db.Set("temp_key", "temp_value", time.Second*2)
	if err != nil {
		t.Fatalf("设置键值对失败: %v", err)
	}

	// 立即获取应该成功
	var value string
	err = db.Get("temp_key", &value)
	if err != nil {
		t.Fatalf("获取键值对失败: %v", err)
	}
	if value != "temp_value" {
		t.Errorf("期望 temp_value, 得到 %s", value)
	}

	// 等待过期
	time.Sleep(time.Second * 3)

	// 过期后获取应该失败
	err = db.Get("temp_key", &value)
	if err == nil {
		t.Error("过期的键应该无法获取")
	}
}

// TestInfo 测试获取数据库信息
func TestInfo(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 添加一些数据
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		t.Fatalf("批量设置失败: %v", err)
	}

	// 获取数据库信息
	info, err := db.Info()
	if err != nil {
		t.Fatalf("获取数据库信息失败: %v", err)
	}

	if info.NumKeys != 2 {
		t.Errorf("期望 2 个键, 得到 %d 个", info.NumKeys)
	}
}

// TestClose 测试关闭数据库
func TestClose(t *testing.T) {
	db := setupTestDB(t)

	err := db.Close()
	if err != nil {
		t.Fatalf("关闭数据库失败: %v", err)
	}

	if !db.IsClosed() {
		t.Error("数据库应该处于关闭状态")
	}

	// 关闭后操作应该失败
	err = db.Set("key", "value", 0)
	if err == nil {
		t.Error("关闭后的数据库不应该允许写入")
	}

	// 清理
	os.RemoveAll(testDBDir)
}

// TestBackupAndRestore 测试备份和恢复
func TestBackupAndRestore(t *testing.T) {
	backupPath := "./testdata/backup.badger"
	os.RemoveAll(backupPath)

	db := setupTestDB(t)

	// 添加一些数据
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	err := db.BatchSet(items, 0)
	if err != nil {
		teardownTestDB(t, db)
		t.Fatalf("批量设置失败: %v", err)
	}

	// 备份
	err = db.Backup(backupPath)
	if err != nil {
		teardownTestDB(t, db)
		t.Fatalf("备份失败: %v", err)
	}

	// 关闭第一个数据库
	teardownTestDB(t, db)

	// 验证备份文件存在
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		t.Fatal("备份文件不存在")
	}

	// 创建新数据库并恢复
	newDB := setupTestDB(t)
	defer teardownTestDB(t, newDB)

	err = newDB.Restore(backupPath)
	if err != nil {
		t.Fatalf("恢复失败: %v", err)
	}

	// 验证恢复的数据
	var value string
	err = newDB.Get("key1", &value)
	if err != nil {
		t.Fatalf("获取恢复的键失败: %v", err)
	}
	if value != "value1" {
		t.Errorf("期望 value1, 得到 %s", value)
	}

	// 清理备份文件
	os.Remove(backupPath)
}
