package main

import (
	"errors"
)

const (
	StatusPending    = "PENDING"
	StatusProcessing = "PROCESSING"
	StatusCompleted  = "COMPLETED"
)

type Book struct {
	ID     string
	Title  string
	Pages  int
	Status string
}

type BookManager struct {
	books map[string]Book
}

func NewBookManager() *BookManager {
	return &BookManager{
		books: make(map[string]Book),
	}
}

func (m *BookManager) Add(id, title string, pages int) error {
	if title == "" || pages <= 0 {
		return errors.New("bad input")
	}

	if _, ok := m.books[id]; ok {
		return errors.New("already exists")
	}

	m.books[id] = Book{
		ID:     id,
		Title:  title,
		Pages:  pages,
		Status: StatusPending,
	}
	return nil
}

func (m *BookManager) Process(id string) error {
	book, ok := m.books[id]
	if !ok {
		return errors.New("not found")
	}

	if book.Status != StatusPending {
		return errors.New("already processed")
	}

	book.Status = StatusProcessing
	m.books[id] = book

	book.Status = StatusCompleted
	m.books[id] = book

	return nil
}
