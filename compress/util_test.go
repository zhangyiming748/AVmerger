package main

import (
	"errors"
	"fmt"
	"testing"
)

// go test -v -run  TestCompress
func TestCompress(t *testing.T) {

}

func anErr() error {
	err := errors.New("it is an err")
	return err
}
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover err:", err)
		} else {
			fmt.Println("no err:", err)
		}
	}()
	anErr()
}
