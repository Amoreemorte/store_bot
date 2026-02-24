package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func GetGorm(cfg *GormConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.GetDCN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect gorm to postgres: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get sql db: %w", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	return db, nil
}
