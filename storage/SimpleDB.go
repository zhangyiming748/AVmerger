package storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SimpleDB struct {
	FilePath string
}

// 创建新的数据库
func NewSimpleDB(filePath string) *SimpleDB {
	return &SimpleDB{FilePath: filePath}
}

// 添加或更新键值对
func (db *SimpleDB) Set(key, value string) error {
	data, err := db.ReadAll()
	if err != nil {
		return err
	}

	// 更新或添加
	data[key] = value
	return db.WriteAll(data)
}

// 查询值
func (db *SimpleDB) Get(key string) (string, bool) {
	data, err := db.ReadAll()
	if err != nil {
		return "", false
	}
	val, exists := data[key]
	return val, exists
}

// 删除键值对
func (db *SimpleDB) Delete(key string) error {
	data, err := db.ReadAll()
	if err != nil {
		return err
	}

	delete(data, key)
	return db.WriteAll(data)
}

// 读取所有数据
func (db *SimpleDB) ReadAll() (map[string]string, error) {
	data := make(map[string]string)
	file, err := os.Open(db.FilePath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			data[parts[0]] = parts[1]
		}
	}
	return data, scanner.Err()
}

// 写入所有数据
func (db *SimpleDB) WriteAll(data map[string]string) error {
	file, err := os.OpenFile(db.FilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range data {
		_, err := fmt.Fprintf(file, "%s=%s\n", k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
