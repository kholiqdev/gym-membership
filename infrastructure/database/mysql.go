package database

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_adminEntity "gym/cmd/domain/admin/entity"
	_classEntity "gym/cmd/domain/class/entity"
	_classBookingEntity "gym/cmd/domain/class_booking/entity"
	_trainerEntity "gym/cmd/domain/class_category/entity"
	_memberEntity "gym/cmd/domain/member/entity"
	"gym/config"
)

type DatabaseImpl struct {
	DB *gorm.DB
}

func NewDatabaseClient(cfg *config.Config, mode string) *DatabaseImpl {
	if &mode == nil {
		mode = "live"
	}

	switch mode {
	case "test":
		return TestDatabase(cfg)
	default:
		return LiveDatabase(cfg)
	}

}

func LiveDatabase(cfg *config.Config) *DatabaseImpl {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	var dbDialector gorm.Dialector

	switch cfg.DBDriver {
	case "postgres":
		dbDialector = postgres.Open(connectionString)
		log.Info().Str("Driver", cfg.DBDriver).Msg("Try to connect with postgres driver")
	default:
		dbDialector = mysql.Open(connectionString)
		log.Info().Str("Driver", cfg.DBDriver).Msg("Try to connect with mysql driver")
	}

	db, err := gorm.Open(dbDialector, &gorm.Config{})

	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	initMigrate(db)

	log.Info().Str("Name", cfg.DBName).Msg("Success connect to DB")
	return &DatabaseImpl{
		DB: db,
	}
}

func TestDatabase(cfg *config.Config) *DatabaseImpl {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	var dbDialector gorm.Dialector

	switch cfg.TestDBDriver {
	case "postgres":
		dbDialector = postgres.Open(connectionString)
		log.Info().Str("Driver", cfg.TestDBDriver).Msg("Try to connect with postgres driver")
	default:
		dbDialector = mysql.Open(connectionString)
		log.Info().Str("Driver", cfg.TestDBDriver).Msg("Try to connect with mysql driver")
	}

	db, err := gorm.Open(dbDialector, &gorm.Config{})

	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	db.Debug()

	initMigrate(db)

	log.Info().Str("Name", cfg.TestDBName).Msg("Success connect to DB")
	return &DatabaseImpl{
		DB: db,
	}
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_adminEntity.Admin{},
		&_memberEntity.Member{},
		&_memberEntity.MemberType{},
		&_memberEntity.MemberJoin{},
		&_trainerEntity.ClassCategory{},
		&_classEntity.Class{},
		&_classBookingEntity.ClassBooking{},
		&_classBookingEntity.ClassBookingDetail{},
	)
}
