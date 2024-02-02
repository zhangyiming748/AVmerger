package sql

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	SetEngine()
	b := new(Bili)
	defer func() {
		if err := recover(); err != nil {
			b.Success = true
		}
		b.SetOne()
	}()
	b.Format = "hevc"
	b.AvID = "312"
	b.BvID = "321"
	b.Cover = "https"
	b.Title = "title"
	b.PartName = "partname"
	b.Success = false
	//panic("故意的")
}

func TestS2t(t *testing.T) {
	S2T("3")
}

func TestTwice(t *testing.T) {
	once()
}

func once() {
	SetEngine()
	b := new(Bili)
	b.Format = "hevc"
	b.AvID = "312"
	b.BvID = "321"
	b.Cover = "https"
	b.Title = "title"
	b.PartName = "partname"
	b.Success = false
	fmt.Printf("once 函数的b=%+v\n地址%v\n", b, &b)
	twice(b)
	fmt.Printf("once 函数被更改后的b=%+v\n地址%v\n", b, &b)

}
func twice(b *Bili) {
	b.PartName = "secondName"
	fmt.Printf("twice函数的b=%+v\n地址%v\n", b, &b)
}
