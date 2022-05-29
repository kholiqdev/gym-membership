package entity

import (
	"database/sql"
	"gorm.io/gorm"
	_eMember "gym/cmd/domain/member/entity"
	_eOnlineClass "gym/cmd/domain/online_class/entity"
	"time"
)

type OnlineClassBooking struct {
	ID            uint `gorm:"primaryKey;autoIncrement;<-:create"`
	MemberID      uint
	InvoiceNo     string `gorm:"not null"`
	MemberName    string `gorm:"not null"`
	MemberEmail   string `gorm:"not null"`
	MemberPhone   string `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	Note          sql.NullString
	Status        uint `gorm:"not null;default:1;comment:'0=Cancel,1=Not Paid,2=Confirmed,3=Done'"`
	Total         float64

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Member                   _eMember.Member             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OnlineClassBookingDetail []*OnlineClassBookingDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OnlineClassBookingList []*OnlineClassBooking

type OnlineClassBookingDetail struct {
	ID                   uint `gorm:"primaryKey;autoIncrement;<-:create"`
	OnlineClassBookingID uint
	OnlineClassID        uint
	ClassName            string  `gorm:"not null"`
	ClassDescription     string  `gorm:"not null"`
	ClassQuota           uint    `gorm:"not null"`
	ClassImage           string  `gorm:"not null"`
	ClassPrice           float64 `gorm:"not null"`
	ClassDuration        uint    `gorm:"not null;default:1;comment:'0=If set to 0,then duration is TBA'"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	OnlineClass        _eOnlineClass.OnlineClass `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OnlineClassBooking OnlineClassBooking        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
