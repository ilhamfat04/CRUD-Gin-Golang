package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}
