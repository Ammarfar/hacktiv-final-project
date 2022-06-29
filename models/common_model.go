package models

import (
	"database/sql"
	"time"
)

type Common struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
}
