package entity

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;<-:create"`
	Name     string `gorm:"type:varchar(255) not null"`
	Phone    string `gorm:"type:varchar(255) not null"`
	Email    string `gorm:"type:varchar(255) unique;<-:create"`
	Password string `gorm:"type:varchar(255) not null"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

type AdminList []*Admin
