package class_category

import (
	"github.com/google/wire"
	_rClassCategory "gym/cmd/domain/class_category/repository"
	_sClassCategory "gym/cmd/domain/class_category/service"
	"gym/infrastructure/database"
	"sync"
)

var (
	hdl     *ClassCategoryHandlerImpl
	hdlOnce sync.Once

	svc     *_sClassCategory.ClassCategoryServiceImpl
	svcOnce sync.Once

	repo     *_rClassCategory.ClassCategoryRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(ClassCategoryHandler), new(*ClassCategoryHandlerImpl)),
		wire.Bind(new(_sClassCategory.ClassCategoryService), new(*_sClassCategory.ClassCategoryServiceImpl)),
		wire.Bind(new(_rClassCategory.ClassCategoryRepository), new(*_rClassCategory.ClassCategoryRepositoryImpl)),
	)
)

func ProvideHandler(svc _sClassCategory.ClassCategoryService) (*ClassCategoryHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &ClassCategoryHandlerImpl{
			SvcClassCategory: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rClassCategory.ClassCategoryRepository) (*_sClassCategory.ClassCategoryServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sClassCategory.ClassCategoryServiceImpl{
			RepoClassCategory: repo,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rClassCategory.ClassCategoryRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rClassCategory.ClassCategoryRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
