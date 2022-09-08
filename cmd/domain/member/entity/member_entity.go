package entity

import (
	"gorm.io/gorm"
	"time"
)

type Member struct {
	ID           uint `gorm:"primaryKey;autoIncrement;<-:create"`
	MemberTypeID *uint
	Name         string `gorm:"type:varchar(50) not null"`
	Phone        string `gorm:"type:varchar(50) not null"`
	Email        string `gorm:"type:varchar(50) unique;<-:create"`
	Password     string `gorm:"type:varchar(255) not null"`
	IsActive     string `gorm:"not null"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	MemberType  MemberType    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MemberOrder []*MemberJoin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type MemberList []*Member

type MemberType struct {
	ID          uint    `gorm:"primaryKey;autoIncrement;<-:create"`
	Name        string  `gorm:"type:varchar(50) not null"`
	Description string  `gorm:"type:text not null"`
	Image       string  `gorm:"not null"`
	Duration    uint    `gorm:"type:tinyint not null;default:1;comment:'1=1Month.2=2Month,3=3Month'"`
	Price       float64 `gorm:"not null"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Member []*Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type MemberJoin struct {
	ID                 uint `gorm:"primaryKey;autoIncrement;<-:create"`
	MemberID           uint
	MemberTypeID       uint
	StartAt            string  `gorm:"type:DATE NOT NULL"`
	InvoiceNo          string  `gorm:"not null"`
	MemberNik          string  `gorm:"not null"`
	MemberName         string  `gorm:"not null"`
	MemberEmail        string  `gorm:"not null"`
	MemberPhone        string  `gorm:"not null"`
	MemberGender       string  `gorm:"not null"`
	MemberAddress      string  `gorm:"not null"`
	MemberCity         string  `gorm:"not null"`
	MemberPostalCode   string  `gorm:"not null"`
	MemberTypeName     string  `gorm:"not null"`
	MemberTypeImage    string  `gorm:"not null"`
	MemberTypeDuration uint    `gorm:"not null"`
	MemberTypePrice    float64 `gorm:"not null"`
	PaymentMethod      string  `gorm:"not null"`
	Status             uint    `gorm:"not null;default:1;comment:'0=Cancel,1=Not Paid,2=Paid'"`
	Total              float64

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Member     Member     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MemberType MemberType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
