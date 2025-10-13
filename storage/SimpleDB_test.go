package storage

import (
	"fmt"
	"testing"
)

func TestSimpleDB(t *testing.T) {
	db := NewSimpleDB("C:\\Users\\zen\\Github\\AVmerger\\storage\\data.txt")
	err := db.Set("name", "Alice")
	if err != nil {
		t.Errorf("Failed to set key-value pair: %v", err)
	}
	err = db.Set("age", "30")
	if err != nil {
		t.Errorf("Failed to set key-value pair: %v", err)
	}

	if name, exists := db.Get("name"); exists {
		fmt.Println("Name:", name)
	}

	err = db.Delete("age")
	if err != nil {
		t.Errorf("Failed to delete key-value pair: %v", err)
	}
}
