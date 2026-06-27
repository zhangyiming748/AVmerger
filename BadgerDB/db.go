package badgerdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
)

// DB 封装 BadgerDB 数据库操作
type DB struct {
	db     *badger.DB
	mu     sync.RWMutex
	closed bool
}

// Config 数据库配置
type Config struct {
	Dir              string // 数据库目录
	ValueLogDir      string // 值日志目录（可选，为空则使用 Dir）
	MemTableSize     int64  // 内存表大小
	ValueLogFileSize int64  // 值日志文件大小
	MaxLevels        int    // LSM 树最大层数
	SyncWrites       bool   // 同步写入
	NumCompactors    int    // 压缩器数量
}

// DefaultConfig 返回默认配置
func DefaultConfig(dir string) Config {
	return Config{
		Dir:              dir,
		ValueLogDir:      "",
		MemTableSize:     64 << 20, // 64MB
		ValueLogFileSize: 1 << 30,  // 1GB
		MaxLevels:        7,
		SyncWrites:       false,
		NumCompactors:    4,
	}
}

// New 创建并初始化数据库实例
func New(config Config) (*DB, error) {
	opts := badger.DefaultOptions(config.Dir)

	if config.ValueLogDir != "" {
		opts = opts.WithValueDir(config.ValueLogDir)
	}

	opts = opts.
		WithMemTableSize(config.MemTableSize).
		WithValueLogFileSize(config.ValueLogFileSize).
		WithMaxLevels(config.MaxLevels).
		WithSyncWrites(config.SyncWrites).
		WithNumCompactors(config.NumCompactors).
		WithLoggingLevel(badger.WARNING)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	log.Printf("数据库初始化成功: %s", config.Dir)

	return &DB{
		db:     db,
		closed: false,
	}, nil
}

// Close 关闭数据库
func (d *DB) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.closed {
		return nil
	}

	if err := d.db.Close(); err != nil {
		return fmt.Errorf("关闭数据库失败: %w", err)
	}

	d.closed = true
	log.Println("数据库已关闭")
	return nil
}

// IsClosed 检查数据库是否已关闭
func (d *DB) IsClosed() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.closed
}

// GetDB 获取底层 badger.DB 实例（谨慎使用）
func (d *DB) GetDB() *badger.DB {
	return d.db
}

// Set 设置键值对
// key: 键
// value: 值（任意类型，会自动序列化为 JSON）
// ttl: 过期时间（0 表示永不过期）
func (d *DB) Set(key string, value interface{}, ttl time.Duration) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("序列化值失败: %w", err)
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	err = d.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), data)
		if ttl > 0 {
			entry = entry.WithTTL(ttl)
		}
		return txn.SetEntry(entry)
	})

	if err != nil {
		return fmt.Errorf("设置键值对失败 [key=%s]: %w", key, err)
	}

	return nil
}

// Get 获取键对应的值
// key: 键
// value: 用于接收值的指针（必须是可反序列化的类型）
// 返回值不存在时返回 ErrKeyNotFound
func (d *DB) Get(key string, value interface{}) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	var data []byte
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return fmt.Errorf("键不存在: %s", key)
			}
			return err
		}

		return item.Value(func(val []byte) error {
			data = append([]byte{}, val...)
			return nil
		})
	})

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, value); err != nil {
		return fmt.Errorf("反序列化值失败 [key=%s]: %w", key, err)
	}

	return nil
}

// Delete 删除键值对
func (d *DB) Delete(key string) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	err := d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})

	if err != nil {
		return fmt.Errorf("删除键值对失败 [key=%s]: %w", key, err)
	}

	return nil
}

// Exists 检查键是否存在
func (d *DB) Exists(key string) (bool, error) {
	if d.IsClosed() {
		return false, fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	err := d.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		return err
	})

	if err == badger.ErrKeyNotFound {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("检查键存在性失败 [key=%s]: %w", key, err)
	}

	return true, nil
}

// BatchSet 批量设置键值对
// items: 键值对映射
// ttl: 过期时间（0 表示永不过期）
func (d *DB) BatchSet(items map[string]interface{}, ttl time.Duration) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	wb := d.db.NewWriteBatch()
	defer wb.Cancel()

	for key, value := range items {
		data, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("序列化值失败 [key=%s]: %w", key, err)
		}

		entry := badger.NewEntry([]byte(key), data)
		if ttl > 0 {
			entry = entry.WithTTL(ttl)
		}

		if err := wb.SetEntry(entry); err != nil {
			return fmt.Errorf("批量设置失败 [key=%s]: %w", key, err)
		}
	}

	if err := wb.Flush(); err != nil {
		return fmt.Errorf("批量刷新失败: %w", err)
	}

	return nil
}

// BatchDelete 批量删除键值对
func (d *DB) BatchDelete(keys []string) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	wb := d.db.NewWriteBatch()
	defer wb.Cancel()

	for _, key := range keys {
		if err := wb.Delete([]byte(key)); err != nil {
			return fmt.Errorf("批量删除失败 [key=%s]: %w", key, err)
		}
	}

	if err := wb.Flush(); err != nil {
		return fmt.Errorf("批量刷新失败: %w", err)
	}

	return nil
}

// GetAllKeys 获取所有键
func (d *DB) GetAllKeys() ([]string, error) {
	if d.IsClosed() {
		return nil, fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	var keys []string
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			keys = append(keys, string(item.Key()))
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("获取所有键失败: %w", err)
	}

	return keys, nil
}

// GetKeysWithPrefix 获取指定前缀的所有键
func (d *DB) GetKeysWithPrefix(prefix string) ([]string, error) {
	if d.IsClosed() {
		return nil, fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	var keys []string
	prefixBytes := []byte(prefix)

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixBytes); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			if !hasPrefix(key, prefixBytes) {
				break
			}
			keys = append(keys, string(key))
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("获取前缀键失败 [prefix=%s]: %w", prefix, err)
	}

	return keys, nil
}

// Count 获取键值对总数
func (d *DB) Count() (int, error) {
	if d.IsClosed() {
		return 0, fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	count := 0
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			count++
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("统计键值对数量失败: %w", err)
	}

	return count, nil
}

// Clear 清空数据库
func (d *DB) Clear() error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// 使用 DropAll 清空所有数据
	if err := d.db.DropAll(); err != nil {
		return fmt.Errorf("清空数据库失败: %w", err)
	}

	log.Println("数据库已清空")
	return nil
}

// Backup 备份数据库到指定文件
func (d *DB) Backup(filePath string) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	file, err := createFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := d.db.Backup(file, 0); err != nil {
		return fmt.Errorf("备份数据库失败: %w", err)
	}

	log.Printf("数据库备份成功: %s", filePath)
	return nil
}

// Restore 从备份文件恢复数据库
func (d *DB) Restore(filePath string) error {
	if d.IsClosed() {
		return fmt.Errorf("数据库已关闭")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	file, err := openFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := d.db.Load(file, 10); err != nil {
		return fmt.Errorf("恢复数据库失败: %w", err)
	}

	log.Printf("数据库恢复成功: %s", filePath)
	return nil
}

// Size 获取数据库大小信息
type Size struct {
	LSMSize  int64  // LSM 树大小
	VLogSize int64  // 值日志大小
	NumKeys  uint64 // 键数量
}

// Info 获取数据库基本信息
func (d *DB) Info() (*Size, error) {
	if d.IsClosed() {
		return nil, fmt.Errorf("数据库已关闭")
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	size := &Size{}
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			size.NumKeys++
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("获取数据库信息失败: %w", err)
	}

	// 获取目录大小
	lsmSize, vlogSize := d.db.Size()
	size.LSMSize = lsmSize
	size.VLogSize = vlogSize

	return size, nil
}

// hasPrefix 检查字节切片是否有指定前缀
func hasPrefix(key, prefix []byte) bool {
	if len(key) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if key[i] != prefix[i] {
			return false
		}
	}
	return true
}

// createFile 创建文件
func createFile(path string) (*os.File, error) {
	return os.Create(path)
}

// openFile 打开文件
func openFile(path string) (*os.File, error) {
	return os.Open(path)
}
