package entities

import (
	"time"
)

type Base struct {
	ID        int64     `json:"-"`
	Status    bool      `gorm:"column:status" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

func (b *Base) PrePersist() {
	b.CreatedAt = time.Now()
}
