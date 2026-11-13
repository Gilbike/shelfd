package collection

import "time"

type UserBook struct {
	UserId         int64     `json:"user_id"`
	BookId         int64     `json:"book_id"`
	CollectionName string    `json:"collection"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
