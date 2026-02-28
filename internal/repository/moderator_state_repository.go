package repository

import (
	"fmt"
	"store_bot/internal/models"

	"gorm.io/gorm"
)

type ModeratorStateRepository struct {
	cfg *ModeratorStateRepositoryConfig
	db  *gorm.DB
}

func NewModeratorStateRepository(cfg *ModeratorStateRepositoryConfig, db *gorm.DB) (*ModeratorStateRepository, error) {
	if cfg.MigrateEnable {
		moderatorState := &models.ModeratorState{}
		if err := db.AutoMigrate(moderatorState); err != nil {
			return nil, fmt.Errorf("Unable to migrate moderator state: %w", err)
		}
	}
	return &ModeratorStateRepository{
		db:  db,
		cfg: cfg,
	}, nil
}

func (m *ModeratorStateRepository) CreateModeratorState(state *models.ModeratorState) (*models.ModeratorState, error) {
	if err := m.db.Model(&models.ModeratorState{}).Create(state).Error; err != nil {
		return nil, fmt.Errorf("%s.CreateModeratorState: %w", m.Name(), err)
	}
	return state, nil
}

func (m *ModeratorStateRepository) UpdateModeratorState(moderatorId int64, state models.ModeratorStateRef) error {
	if err := m.db.Model(&models.ModeratorState{}).
		Where("moderatorId = ?", moderatorId).
		Update("state = ?", state).Error; err != nil {
		return fmt.Errorf("%s.UpdateModeratorState: %w", m.Name(), err)
	}
	return nil
}

func (m *ModeratorStateRepository) GetModeratorState(moderatorId int64) (*models.ModeratorState, error) {
	state := &models.ModeratorState{ModeratorId: moderatorId}
	if err := m.db.Model(&models.ModeratorState{}).First(state, moderatorId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetModerator: %w", m.Name(), err)
	}
	return state, nil
}

// deleteUnscoped = false will be used soft deleting
//
// deleteUnscoped = true will be used full deleting
func (m *ModeratorStateRepository) DeleteModeratorState(moderatorId int64, deleteUnscoped bool) error {
	var err error
	if deleteUnscoped {
		err = m.db.Model(&models.ModeratorState{}).Unscoped().Delete("moderatorId = ?", moderatorId).Error
	} else {
		err = m.db.Model(&models.ModeratorState{}).Delete("moderatorId = ?", moderatorId).Error
	}
	if err != nil {
		return fmt.Errorf("%s.DeleteModeratorState: %w", m.Name(), err)
	}
	return nil
}

func (m *ModeratorStateRepository) Name() string {
	return "ModeratorStateRepository"
}
