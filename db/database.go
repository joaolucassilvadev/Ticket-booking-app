package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(DBMigrator func(db *gorm.DB) error) *gorm.DB {
	dsn := "user=postgres.ivdjytiqqebwmuuqvbcr password=evilasio22 host=aws-0-sa-east-1.pooler.supabase.com port=6543 dbname=postgres sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
