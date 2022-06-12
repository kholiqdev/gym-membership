package class

import (
	"github.com/google/wire"
	_rClass "gym/cmd/domain/class/repository"
	_sClass "gym/cmd/domain/class/service"
	"gym/infrastructure/database"
	"sync"
)

var (
	hdl     *ClassHandlerImpl
	hdlOnce sync.Once

	svc     *_sClass.ClassServiceImpl
	svcOnce sync.Once

	repo     *_rClass.ClassRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(ClassHandler), new(*ClassHandlerImpl)),
		wire.Bind(new(_sClass.ClassService), new(*_sClass.ClassServiceImpl)),
		wire.Bind(new(_rClass.ClassRepository), new(*_rClass.ClassRepositoryImpl)),
	)
)

func ProvideHandler(svc _sClass.ClassService) (*ClassHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &ClassHandlerImpl{
			SvcClass: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rClass.ClassRepository) (*_sClass.ClassServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sClass.ClassServiceImpl{
			RepoClass: repo,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rClass.ClassRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rClass.ClassRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
