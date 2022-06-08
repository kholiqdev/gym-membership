package repository

import (
	"gorm.io/gorm"
	"gym/cmd/domain/member/entity"
)

type MemberRepositoryImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	Db *gorm.DB
}

func (r *MemberRepositoryImpl) FindAll() (*entity.MemberList, error) {
	var members entity.MemberList

	if e := r.Db.Debug().Find(&members).Error; e != nil {
		return nil, e
	}

	return &members, nil
}

func (r *MemberRepositoryImpl) Find(memberId uint) (*entity.Member, error) {
	var member entity.Member

	if e := r.Db.Debug().First(&member, memberId).Error; e != nil {
		return nil, e
	}

	return &member, nil
}

func (r *MemberRepositoryImpl) FindByEmail(email string) (*entity.Member, error) {
	var member entity.Member

	if e := r.Db.Debug().Where("email = ?", email).First(&member).Error; e != nil {
		return nil, e
	}

	return &member, nil
}

func (r *MemberRepositoryImpl) Insert(member *entity.Member) (*entity.Member, error) {
	if e := r.Db.Debug().Create(&member).Error; e != nil {
		return nil, e
	}
	return member, nil
}

func (r *MemberRepositoryImpl) InsertMemberType(memberType *entity.MemberType) (*entity.MemberType, error) {
	if e := r.Db.Debug().Create(&memberType).Error; e != nil {
		return nil, e
	}
	return memberType, nil
}
