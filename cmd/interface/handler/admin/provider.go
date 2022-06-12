package admin

import (
	"github.com/google/wire"
	_rAdmin "gym/cmd/domain/admin/repository"
	_sAdmin "gym/cmd/domain/admin/service"
	_sClass "gym/cmd/domain/class/service"
	_sClassCategory "gym/cmd/domain/class_category/service"
	_sMember "gym/cmd/domain/member/service"
	"gym/infrastructure/database"
	"gym/pkg/auth"
	"sync"
)

var (
	hdl     *AdminHandlerImpl
	hdlOnce sync.Once

	svc     *_sAdmin.AdminServiceImpl
	svcOnce sync.Once

	repo     *_rAdmin.AdminRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(AdminHandler), new(*AdminHandlerImpl)),
		wire.Bind(new(_sAdmin.AdminService), new(*_sAdmin.AdminServiceImpl)),
		wire.Bind(new(_rAdmin.AdminRepository), new(*_rAdmin.AdminRepositoryImpl)),
	)
)

func ProvideHandler(svc _sAdmin.AdminService, svcMember _sMember.MemberService, svcClass _sClass.ClassService, svcClassCategory _sClassCategory.ClassCategoryService) (*AdminHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &AdminHandlerImpl{
			SvcAdmin:         svc,
			SvcMember:        svcMember,
			SvcClass:         svcClass,
			SvcClassCategory: svcClassCategory,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rAdmin.AdminRepository, jwtAuth auth.JwtToken) (*_sAdmin.AdminServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sAdmin.AdminServiceImpl{
			RepoAdmin: repo,
			JwtAuth:   jwtAuth,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rAdmin.AdminRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rAdmin.AdminRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
