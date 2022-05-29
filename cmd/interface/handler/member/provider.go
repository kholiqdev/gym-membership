package member

import (
	"github.com/google/wire"
	_rMember "gym/cmd/domain/member/repository"
	_sMember "gym/cmd/domain/member/service"
	"gym/infrastructure/database"
	"gym/pkg/auth"
	"sync"
)

var (
	hdl     *MemberHandlerImpl
	hdlOnce sync.Once

	svc     *_sMember.MemberServiceImpl
	svcOnce sync.Once

	repo     *_rMember.MemberRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(MemberHandler), new(*MemberHandlerImpl)),
		wire.Bind(new(_sMember.MemberService), new(*_sMember.MemberServiceImpl)),
		wire.Bind(new(_rMember.MemberRepository), new(*_rMember.MemberRepositoryImpl)),
	)
)

func ProvideHandler(svc _sMember.MemberService) (*MemberHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &MemberHandlerImpl{
			SvcMember: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rMember.MemberRepository, jwtAuth auth.JwtToken) (*_sMember.MemberServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sMember.MemberServiceImpl{
			RepoMember: repo,
			JwtAuth:    jwtAuth,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rMember.MemberRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rMember.MemberRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
