package entity

import (
	"gorm.io/gorm"
	"time"
)

type ClassCategory struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;<-:create"`
	Name string `gorm:"type:varchar(255) not null"`
	Slug string `gorm:"unique;not null;size:50"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

type ClassCategoryList []*ClassCategory
