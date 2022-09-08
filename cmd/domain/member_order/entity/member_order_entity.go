package entity

import (
	"database/sql"
	"gorm.io/gorm"
	_eClass "gym/cmd/domain/class/entity"
	_eMember "gym/cmd/domain/member/entity"
	"time"
)

type MemberOrder struct {
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
	Member            _eMember.Member      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MemberOrderDetail []*MemberOrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type MemberOrderList []*MemberOrder

type MemberOrderDetail struct {
	ID               uint `gorm:"primaryKey;autoIncrement;<-:create"`
	MemberOrderID    uint
	ClassID          uint
	ClassName        string `gorm:"not null"`
	ClassDescription string `gorm:"not null"`
	ClassMeetLink    sql.NullString
	ClassCategory    uint      `gorm:"not null"`
	ClassImage       string    `gorm:"not null"`
	ClassPrice       float64   `gorm:"not null"`
	ClassDate        time.Time `gorm:"type:DATE NOT NULL"`
	ClassStartTime   time.Time `gorm:"type:TIME NOT NULL"`
	ClassEndTime     time.Time `gorm:"type:TIME NOT NULL"`
	ClassTrainerName string    `gorm:"not null"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Class       _eClass.Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MemberOrder MemberOrder   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
