package repository

import (
	"fmt"
	"store_bot/internal/models"

	"gorm.io/gorm"
)

type CardRepository struct {
	cfg *CardRepositoryConfig
	db  *gorm.DB
}

func NewCardRepository(cfg *CardRepositoryConfig, db *gorm.DB) (*CardRepository, error) {
	if cfg.MigrateEnable {
		card := &models.Card{}
		if err := db.AutoMigrate(card); err != nil {
			return nil, fmt.Errorf("Unable to migrate card: %w", err)
		}
	}
	return &CardRepository{
		db:  db,
		cfg: cfg,
	}, nil
}

func (c *CardRepository) Name() string {
	return "CardRepository"
}

func (c *CardRepository) CreateCard(Card *models.Card) (*models.Card, error) {
	if err := c.db.Model(&models.Card{}).Create(Card).Error; err != nil {
		return nil, fmt.Errorf("%s.CreateCard: %w", c.Name(), err)
	}
	return Card, nil
}

func (c *CardRepository) GetCard(moderatorId int64) (*models.Card, error) {
	Card := &models.Card{ModeratorId: moderatorId}
	if err := c.db.Model(&models.Card{}).Where("ModeratorId = ? AND IsFinished=false", moderatorId).First(Card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetCard: %w", c.Name(), err)
	}
	return Card, nil
}

func (c *CardRepository) UpdateCard(card *models.Card) error {
	if err := c.db.Model(&models.Card{}).
		Where("name = ?", card.Name).
		Updates(card).Error; err != nil {
		return fmt.Errorf("%s.UpdateCard: %w", c.Name(), err)
	}
	return nil
}

// deleteUnscoped = false will be used soft deleting
//
// deleteUnscoped = true will be used full deleting
func (m *CardRepository) DeleteCard(name string, deleteUnscoped bool) error {
	var err error
	if deleteUnscoped {
		err = m.db.Unscoped().Where("name = ?", name).Delete(&models.Card{}).Error
	} else {
		err = m.db.Where("name = ?", name).Delete(&models.Card{}).Error
	}
	if err != nil {
		return fmt.Errorf("%s.DeleteCard: %w", m.Name(), err)
	}
	return nil
}
