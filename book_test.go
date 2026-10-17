package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewBookManager()

	err := m.Add("1", "Go Book", 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	book := m.books["1"]
	if book.Title != "Go Book" {
		t.Errorf("expected title Go Book, got %s", book.Title)
	}
}

func TestAddError(t *testing.T) {
	m := NewBookManager()

	err := m.Add("1", "", 0)
	if err == nil {
		t.Error("expected error for empty title and 0 pages, got nil")
	}
}
