package repository

import (
	"fmt"
	"store_bot/internal/models"
	handlers "store_bot/internal/services"

	"gorm.io/gorm"
)

type ModeratorRepository struct {
	cfg *ModeratorRepositoryConfig
	db  *gorm.DB
}

var _ handlers.ModeratorRepository = (*ModeratorRepository)(nil)

func NewModeratorRepository(cfg *ModeratorRepositoryConfig, db *gorm.DB) (*ModeratorRepository, error) {
	if cfg.MigrateEnable {
		moderator := &models.Moderator{}
		if err := db.AutoMigrate(moderator); err != nil {
			return nil, fmt.Errorf("Unable to migrate moderator: %w", err)
		}
	}
	return &ModeratorRepository{
		db:  db,
		cfg: cfg,
	}, nil
}

func (m *ModeratorRepository) CreateModerator(moderator *models.Moderator) (*models.Moderator, error) {
	if err := m.db.Model(&models.Moderator{}).Create(moderator).Error; err != nil {
		return nil, fmt.Errorf("%s.CreateModerator: %w", m.Name(), err)
	}
	return moderator, nil
}

func (m *ModeratorRepository) GetModerator(id int64) (*models.Moderator, error) {
	moderator := &models.Moderator{ID: id}
	if err := m.db.Model(&models.Moderator{}).Where(moderator).First(moderator).Error; err != nil {
		return nil, fmt.Errorf("%s.GetModerator: %w", m.Name(), err)
	}
	return moderator, nil
}

// deleteUnscoped = false will be used soft deleting
//
// deleteUnscoped = true will be used full deleting
func (m *ModeratorRepository) DeleteModerator(id int64, deleteUnscoped bool) error {
	var err error
	if deleteUnscoped {
		err = m.db.Model(&models.Moderator{}).Unscoped().Delete("id = ?", id).Error
	} else {
		err = m.db.Model(&models.Moderator{}).Delete("id = ?", id).Error
	}
	if err != nil {
		return fmt.Errorf("%s.DeleteModerator: %w", m.Name(), err)
	}
	return nil
}

func (m *ModeratorRepository) Name() string {
	return "ModeratorRepository"
}
