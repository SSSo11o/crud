package models

import "time"

type Url struct {
	ID        int64     `gorm:"primary key; serial" json:"id"`
	Name      string    `gorm:"type: text;not null" json:"name"`
	Method    string    `gorm:"type: text;not null" json:"method"`
	IsActive  bool      `gorm:"type: boolean;default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Url) TableName() string { return "url" }
