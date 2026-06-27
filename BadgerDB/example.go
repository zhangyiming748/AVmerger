package badgerdb

import (
	"fmt"
	"log"
	"time"
)

// Example 演示如何使用 BadgerDB 封装
func Example() {
	// 1. 初始化数据库
	config := DefaultConfig("./data/badgerdb")
	db, err := New(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 设置单个键值对
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	user := User{ID: 1, Name: "张三", Email: "zhangsan@example.com"}
	err = db.Set("user:1", user, 0) // 0 表示永不过期
	if err != nil {
		log.Printf("设置失败: %v", err)
	}

	// 3. 设置带过期时间的键值对
	err = db.Set("session:abc123", map[string]string{"token": "xyz"}, time.Hour*24)
	if err != nil {
		log.Printf("设置失败: %v", err)
	}

	// 4. 获取值
	var retrievedUser User
	err = db.Get("user:1", &retrievedUser)
	if err != nil {
		log.Printf("获取失败: %v", err)
	} else {
		fmt.Printf("用户信息: %+v\n", retrievedUser)
	}

	// 5. 检查键是否存在
	exists, err := db.Exists("user:1")
	if err != nil {
		log.Printf("检查失败: %v", err)
	} else {
		fmt.Printf("user:1 是否存在: %v\n", exists)
	}

	// 6. 批量设置
	users := map[string]interface{}{
		"user:2": User{ID: 2, Name: "李四", Email: "lisi@example.com"},
		"user:3": User{ID: 3, Name: "王五", Email: "wangwu@example.com"},
	}
	err = db.BatchSet(users, 0)
	if err != nil {
		log.Printf("批量设置失败: %v", err)
	}

	// 7. 获取所有键
	keys, err := db.GetAllKeys()
	if err != nil {
		log.Printf("获取所有键失败: %v", err)
	} else {
		fmt.Printf("所有键: %v\n", keys)
	}

	// 8. 获取指定前缀的键
	userKeys, err := db.GetKeysWithPrefix("user:")
	if err != nil {
		log.Printf("获取前缀键失败: %v", err)
	} else {
		fmt.Printf("用户键: %v\n", userKeys)
	}

	// 9. 统计数量
	count, err := db.Count()
	if err != nil {
		log.Printf("统计失败: %v", err)
	} else {
		fmt.Printf("总键数: %d\n", count)
	}

	// 10. 获取数据库信息
	info, err := db.Info()
	if err != nil {
		log.Printf("获取信息失败: %v", err)
	} else {
		fmt.Printf("数据库大小 - LSM: %d bytes, VLog: %d bytes, Keys: %d\n",
			info.LSMSize, info.VLogSize, info.NumKeys)
	}

	// 11. 删除单个键
	err = db.Delete("session:abc123")
	if err != nil {
		log.Printf("删除失败: %v", err)
	}

	// 12. 批量删除
	err = db.BatchDelete([]string{"user:2", "user:3"})
	if err != nil {
		log.Printf("批量删除失败: %v", err)
	}

	// 13. 备份数据库
	err = db.Backup("./data/backup.badger")
	if err != nil {
		log.Printf("备份失败: %v", err)
	}

	// 14. 清空数据库（谨慎使用）
	// err = db.Clear()
	// if err != nil {
	// 	log.Printf("清空失败: %v", err)
	// }

	fmt.Println("示例完成!")
}
