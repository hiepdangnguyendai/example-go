package domain

import (
	"time"
)

type Lend struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"book_id"`
	From   time.Time `sql:"default:now()" json:"from"`
	To     time.Time `json:"to"`
}
