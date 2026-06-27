# BadgerDB 封装模块

这是一个基于 BadgerDB v4 的键值数据库封装模块，提供了简单易用的 API。

## 功能特性

- ✅ 数据库初始化和配置
- ✅ 基本的增删改查操作
- ✅ 批量操作（批量设置、批量删除）
- ✅ 键前缀查询
- ✅ TTL（过期时间）支持
- ✅ 数据库备份和恢复
- ✅ 线程安全（使用读写锁）
- ✅ 自动 JSON 序列化/反序列化

## 快速开始

### 1. 初始化数据库

```go
import "AVmerger/BadgerDB"

// 使用默认配置
config := badgerdb.DefaultConfig("./data/mydb")
db, err := badgerdb.New(config)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### 2. 自定义配置

```go
config := badgerdb.Config{
    Dir:              "./data/mydb",
    ValueLogDir:      "",                    // 可选，为空则使用 Dir
    MemTableSize:     64 << 20,             // 64MB
    ValueLogFileSize: 1 << 30,              // 1GB
    MaxLevels:        7,
    SyncWrites:       false,
    NumCompactors:    4,
}
db, err := badgerdb.New(config)
```

### 3. 基本操作

#### 设置键值对

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{ID: 1, Name: "张三", Email: "zhangsan@example.com"}

// 永久存储
err := db.Set("user:1", user, 0)

// 带过期时间（24小时后过期）
err := db.Set("session:abc", tokenData, time.Hour*24)
```

#### 获取值

```go
var user User
err := db.Get("user:1", &user)
if err != nil {
    log.Printf("获取失败: %v", err)
}
```

#### 删除键

```go
err := db.Delete("user:1")
```

#### 检查键是否存在

```go
exists, err := db.Exists("user:1")
if err != nil {
    log.Printf("检查失败: %v", err)
}
```

### 4. 批量操作

#### 批量设置

```go
users := map[string]interface{}{
    "user:1": User{ID: 1, Name: "张三"},
    "user:2": User{ID: 2, Name: "李四"},
    "user:3": User{ID: 3, Name: "王五"},
}
err := db.BatchSet(users, 0)
```

#### 批量删除

```go
keys := []string{"user:1", "user:2"}
err := db.BatchDelete(keys)
```

### 5. 查询操作

#### 获取所有键

```go
keys, err := db.GetAllKeys()
```

#### 获取指定前缀的键

```go
userKeys, err := db.GetKeysWithPrefix("user:")
```

#### 统计键数量

```go
count, err := db.Count()
```

### 6. 数据库管理

#### 获取数据库信息

```go
info, err := db.Info()
if err != nil {
    log.Printf("获取信息失败: %v", err)
}
fmt.Printf("LSM大小: %d bytes\n", info.LSMSize)
fmt.Printf("VLog大小: %d bytes\n", info.VLogSize)
fmt.Printf("键数量: %d\n", info.NumKeys)
```

#### 备份数据库

```go
err := db.Backup("./backup/mydb.badger")
```

#### 恢复数据库

```go
err := db.Restore("./backup/mydb.badger")
```

#### 清空数据库

```go
err := db.Clear()
```

## 完整示例

查看 [example.go](example.go) 文件了解完整的使用示例。

运行示例：

```go
package main

import (
    "AVmerger/BadgerDB"
)

func main() {
    badgerdb.Example()
}
```

## 运行测试

```bash
cd BadgerDB
go test -v
```

## 注意事项

1. **线程安全**：所有操作都是线程安全的，可以在多个 goroutine 中并发使用
2. **JSON 序列化**：所有值都会自动序列化为 JSON 存储，确保你的结构体有正确的 JSON 标签
3. **资源释放**：使用完毕后务必调用 `Close()` 方法关闭数据库
4. **键命名**：建议使用冒号分隔的命名方式（如 `user:1`, `session:abc`）便于分类管理
5. **TTL**：设置 TTL 为 0 表示永不过期，大于 0 的值会在指定时间后自动删除

## API 参考

### 核心类型

- `DB`: 数据库实例
- `Config`: 数据库配置
- `Size`: 数据库大小信息

### 主要方法

| 方法 | 说明 |
|------|------|
| `New(config)` | 创建并初始化数据库 |
| `Close()` | 关闭数据库 |
| `Set(key, value, ttl)` | 设置键值对 |
| `Get(key, &value)` | 获取值 |
| `Delete(key)` | 删除键 |
| `Exists(key)` | 检查键是否存在 |
| `BatchSet(items, ttl)` | 批量设置 |
| `BatchDelete(keys)` | 批量删除 |
| `GetAllKeys()` | 获取所有键 |
| `GetKeysWithPrefix(prefix)` | 获取指定前缀的键 |
| `Count()` | 统计键数量 |
| `Info()` | 获取数据库信息 |
| `Backup(path)` | 备份数据库 |
| `Restore(path)` | 恢复数据库 |
| `Clear()` | 清空数据库 |

## 依赖

- [github.com/dgraph-io/badger/v4](https://github.com/dgraph-io/badger)
