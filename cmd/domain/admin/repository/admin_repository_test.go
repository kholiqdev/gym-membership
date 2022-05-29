package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gym/cmd/domain/admin/entity"
	"regexp"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func InitConnection() (sqlmock.Sqlmock, AdminRepositoryImpl) {
	sqlMockDB, sqlMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlMockDB,
		SkipInitializeWithVersion: true,
	}))

	adminRepo := AdminRepositoryImpl{
		Db: gormDB,
	}
	return sqlMock, adminRepo
}

func TestAdminRepository_FindAllFail(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `admins` WHERE `admins`.`deleted_at` IS NULL").
		WillReturnError(errors.New("can't fetch to mock db"))

	admins, err := adminRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, admins)
}

func TestAdminRepository_FindAllSuccess(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `admins` WHERE `admins`.`deleted_at` IS NULL").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "email", "password", "created_at", "updated_at"}))
	admins, err := adminRepo.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, admins)
}

func TestAdminRepository_FindFail(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	sqlMock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `admins` WHERE `admins`.`id` = ? AND `admins`.`deleted_at` IS NULL ORDER BY `admins`.`id` LIMIT 1")).
		WillReturnError(errors.New("can't fetch to mock db"))
	admins, err := adminRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, admins)
}

func TestAdminRepository_FindSuccess(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		phone     string    = "0867676767"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)
	sqlMock.ExpectQuery("SELECT * FROM `admins` WHERE `admins`.`id` = ? AND `admins`.`deleted_at` IS NULL ORDER BY `admins`.`id` LIMIT 1").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "email", "password", "created_at", "updated_at"}).AddRow(id, name, phone, email, password, createdAt, updatedAt))
	admins, err := adminRepo.Find(id)

	assert.Nil(t, err)
	assert.NotNil(t, admins)
	assert.IsType(t, &entity.Admin{}, admins)
}

func TestAdminRepository_FindByEmailFail(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `admins` WHERE email = ? AND `admins`.`deleted_at` IS NULL ORDER BY `admins`.`id` LIMIT 1").
		WillReturnError(errors.New("can't fetch to mock db"))

	admins, err := adminRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, admins)
}

func TestAdminRepository_FindByEmailSuccess(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		phone     string    = "0867676767"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)

	sqlMock.ExpectQuery("SELECT * FROM `admins` WHERE email = ? AND `admins`.`deleted_at` IS NULL ORDER BY `admins`.`id` LIMIT 1").WithArgs(email).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "email", "password", "created_at", "updated_at"}).AddRow(id, name, phone, email, password, createdAt, updatedAt))
	admins, err := adminRepo.FindByEmail(email)

	assert.Nil(t, err)
	assert.NotNil(t, admins)
	assert.IsType(t, &entity.Admin{}, admins)
}

func TestAdminRepository_InsertFail(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		phone     string  = "0867676767"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectExec("INSERT INTO `admins` (`name`,`phone`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)").WithArgs(name, phone, email, password, createdAt, updatedAt, nil).
		WillReturnError(errors.New("can't fetch to mock db"))

	adminEntity := entity.Admin{
		Name:     "Abdul Kholiq",
		Phone:    "0867676767",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	admins, err := adminRepo.Insert(&adminEntity)
	assert.NotNil(t, err)
	assert.Nil(t, admins)
}

func TestAdminRepository_InsertSuccess(t *testing.T) {
	sqlMock, adminRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		phone     string  = "0867676767"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO `admins` (`name`,`phone`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)").
		WithArgs(name, phone, email, password, createdAt, updatedAt, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	adminEntity := entity.Admin{
		Name:     "Abdul Kholiq",
		Phone:    "0867676767",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}
	admin, err := adminRepo.Insert(&adminEntity)

	assert.Nil(t, err)
	assert.NotNil(t, admin)
	assert.IsType(t, &entity.Admin{}, admin)
	assert.Equal(t, email, admin.Email)
}
