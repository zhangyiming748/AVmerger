package test

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("panic之前")
		panic("")
	}()
	wg.Wait()
	t.Log("panic之后")
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保程序退出时取消上下文，防止资源泄露
	go NumsOfGoroutine(ctx)
	for i := 0; i < 10; i++ {
		t.Log("主函数运行")
		time.Sleep(time.Second)
	}
	ctx.Done()
}
func NumsOfGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("NumsOfGoroutine exiting...")
			return
		default:
			fmt.Printf("当前程序运行时协程个数:%d\n", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}
}
