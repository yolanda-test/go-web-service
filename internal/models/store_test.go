package models

import (
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()

	// Test adding an item
	item := "test item"
	err := store.Add(item)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test retrieving the item
	retrievedItem, err := store.Get(item)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if retrievedItem != item {
		t.Fatalf("expected %v, got %v", item, retrievedItem)
	}

	// Test retrieving a non-existent item
	_, err = store.Get("non-existent item")
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
