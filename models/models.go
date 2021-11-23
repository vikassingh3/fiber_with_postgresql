package models

import (
	"github.com/google/uuid"
)

// book struct to describe book object

type Book struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	BookStatus int       `json:"book_status"`
}
