package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"gym/cmd/domain/class_category/entity"
	"time"
)

type Class struct {
	ID              uint `gorm:"primaryKey;autoIncrement;<-:create"`
	ClassCategoryID uint
	TrainerName     string         `gorm:"type:varchar(255) not null"`
	Name            string         `gorm:"type:varchar(255) not null"`
	Description     string         `gorm:"type:varchar(255) not null"`
	Image           string         `gorm:"type:varchar(255) not null"`
	MeetLink        sql.NullString `gorm:"not null"`
	Price           float64        `gorm:"not null"`
	Date            string         `gorm:"type:DATE NOT NULL"`
	StartTime       string         `gorm:"type:TIME NOT NULL"`
	EndTime         string         `gorm:"type:TIME NOT NULL"`
	IsOffline       bool           `gorm:"not null;default:0"`
	CategoryName    string         `gorm:"-:migration;<-:false"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	ClassCategory entity.ClassCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *Class) AfterFind(tx *gorm.DB) (err error) {
	u.CategoryName = u.ClassCategory.Name
	return
}

type ClassList []*Class
