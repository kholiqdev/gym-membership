package entity

import (
	"gorm.io/gorm"
	_trainerEntity "gym/cmd/domain/trainer/entity"
	"time"
)

type OfflineClass struct {
	ID          uint `gorm:"primaryKey;autoIncrement;<-:create"`
	TrainerID   uint
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Quota       uint      `gorm:"not null"`
	Image       string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Date        time.Time `gorm:"type:DATE NOT NULL"`
	Time        time.Time `gorm:"type:TIME NOT NULL"`
	Duration    uint      `gorm:"not null;default:1;comment:'0=If set to 0,then duration is TBA'"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Trainer _trainerEntity.Trainer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OfflineClassList []*OfflineClass
