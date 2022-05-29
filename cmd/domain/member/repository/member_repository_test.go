package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gym/cmd/domain/member/entity"
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

func InitConnection() (sqlmock.Sqlmock, MemberRepositoryImpl) {
	sqlMockDB, sqlMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlMockDB,
		SkipInitializeWithVersion: true,
	}))

	memberRepo := MemberRepositoryImpl{
		Db: gormDB,
	}
	return sqlMock, memberRepo
}

func TestMemberRepository_FindAllFail(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `members` WHERE `members`.`deleted_at` IS NULL").
		WillReturnError(errors.New("can't fetch to mock db"))

	members, err := memberRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, members)
}

func TestMemberRepository_FindAllSuccess(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `members` WHERE `members`.`deleted_at` IS NULL").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}))
	members, err := memberRepo.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, members)
}

func TestMemberRepository_FindFail(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	sqlMock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `members` WHERE `members`.`id` = ? AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1")).
		WillReturnError(errors.New("can't fetch to mock db"))
	members, err := memberRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, members)
}

func TestMemberRepository_FindSuccess(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)
	sqlMock.ExpectQuery("SELECT * FROM `members` WHERE `members`.`id` = ? AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).AddRow(id, name, email, password, createdAt, updatedAt))
	members, err := memberRepo.Find(id)

	assert.Nil(t, err)
	assert.NotNil(t, members)
	assert.IsType(t, &entity.Member{}, members)
}

func TestMemberRepository_FindByEmailFail(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `members` WHERE email = ? AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1").
		WillReturnError(errors.New("can't fetch to mock db"))

	members, err := memberRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, members)
}

func TestMemberRepository_FindByEmailSuccess(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)

	sqlMock.ExpectQuery("SELECT * FROM `members` WHERE email = ? AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1").WithArgs(email).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).AddRow(id, name, email, password, createdAt, updatedAt))
	members, err := memberRepo.FindByEmail(email)

	assert.Nil(t, err)
	assert.NotNil(t, members)
	assert.IsType(t, &entity.Member{}, members)
}

func TestMemberRepository_InsertFail(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		phone     string  = "0859591194"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectExec("INSERT INTO `members` (`name`,`phone`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)").WithArgs(name, phone, email, password, createdAt, updatedAt, nil).
		WillReturnError(errors.New("can't fetch to mock db"))

	memberEntity := entity.Member{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	members, err := memberRepo.Insert(&memberEntity)
	assert.NotNil(t, err)
	assert.Nil(t, members)
}

func TestMemberRepository_InsertSuccess(t *testing.T) {
	sqlMock, memberRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		phone     string  = "0869696969"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO `members` (`name`,`phone`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)").
		WithArgs(name, phone, email, password, createdAt, updatedAt, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	memberEntity := entity.Member{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}
	member, err := memberRepo.Insert(&memberEntity)

	assert.Nil(t, err)
	assert.NotNil(t, member)
	assert.IsType(t, &entity.Member{}, member)
	assert.Equal(t, email, member.Email)
}
