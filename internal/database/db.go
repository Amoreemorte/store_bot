package database

import (
	"fmt"

	"gorm.io/gorm"
)

// Return GORM connected to chosen databse
func GetGorm(cfg *GormConfig, db gorm.Dialector) (*gorm.DB, error) {
	gormDb, err := gorm.Open(db, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect gorm to database: %w", err)
	}
	sqlDB, err := gormDb.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get sql db: %w", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	return gormDb, nil
}
