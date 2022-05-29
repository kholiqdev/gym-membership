package entity

import (
	"gorm.io/gorm"
	_trainerEntity "gym/cmd/domain/trainer/entity"
	"time"
)

type OnlineClass struct {
	ID          uint `gorm:"primaryKey;autoIncrement;<-:create"`
	TrainerID   uint
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Quota       uint    `gorm:"not null"`
	Image       string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Duration    uint    `gorm:"not null;default:1;comment:'0=If set 0(zero),then duration is TBA, the duration in days'"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Trainer _trainerEntity.Trainer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OnlineClassList []*OnlineClass
