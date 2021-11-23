package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type TestModel struct {
	gorm.Model
	Matcher string `gorm:"uniqueIndex;size:512"`
	Data    string
}

func doUpsert(db *gorm.DB) error {
	db.AutoMigrate(&TestModel{})

	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "matcher"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"data": "foobaz"}),
	}).Create(&TestModel{
		Matcher: "foobar",
		Data:    "foobar",
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func main() {
	dsn := os.Getenv("SQLSERVER_DSN")

	// SQLite
	sqliteDb, err := gorm.Open(sqlite.Open("file:gorm.db?mode=memory&cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("sqlite connection: %v", err)
	}
	err = doUpsert(sqliteDb)
	if err != nil {
		log.Fatalf("sqlite upsert: %v", err)
	}

	// SQL Server
	sqlServerDb, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("sql server connection: %v", err)
	}
	err = doUpsert(sqlServerDb)
	if err != nil {
		log.Fatalf("sql server upsert: %v", err)
	}
}
