package repository

import (
	"gym/cmd/domain/member/entity"
)

type MemberRepository interface {
	FindAll() (*entity.MemberList, error)
	Find(memberId uint) (*entity.Member, error)
	FindByEmail(email string) (*entity.Member, error)
	Insert(member *entity.Member) (*entity.Member, error)
	InsertMemberType(member *entity.MemberType) (*entity.MemberType, error)
}
