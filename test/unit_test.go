package test

import (
	"sync"
	"testing"
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
