package database

import (
	"go-codebase/domain/book"
	"go-codebase/domain/users"
	"go-codebase/infrastructure/config"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
)

func Initialize() (*gorm.DB, error) {
	log.Println("Opening connection to database...")
	db, err := gorm.Open(postgres.Open(config.GlobalConfig.Database.URI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return db, err
	}
	return db, nil
}

func MigrateIfNeed(db *gorm.DB) error {
	log.Println("Running database migration if necessary...")
	err := db.AutoMigrate(&users.Users{}, &book.Books{})
	if err != nil {
		return err
	}
	return nil
}
