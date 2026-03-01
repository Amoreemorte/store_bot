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
		moderator := &models.Card{}
		if err := db.AutoMigrate(moderator); err != nil {
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

func (c *CardRepository) CreateCard(card *models.Card) (*models.Card, error) {
	if err := c.db.Model(&models.Card{}).Create(card).Error; err != nil {
		return nil, fmt.Errorf("%s.CreateCard: %w", c.Name(), err)
	}
	return card, nil
}

func (c *CardRepository) GetCard(name string) (*models.Card, error) {
	card := &models.Card{Name: name}
	if err := c.db.Model(&models.Card{}).Where(card).First(card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetCard: %w", c.Name(), err)
	}
	return card, nil
}

func (c *CardRepository) GetCards() ([]models.Card, error) {
	cards := make([]models.Card, 0)
	if err := c.db.Model(&models.Card{}).Find(&cards).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetCards: %w", c.Name(), err)
	}
	return cards, nil
}

// deleteUnscoped = false will be used soft deleting
//
// deleteUnscoped = true will be used full deleting
func (m *CardRepository) DeleteCard(name string, deleteUnscoped bool) error {
	var err error
	if deleteUnscoped {
		err = m.db.Model(&models.Card{}).Unscoped().Where("name = ?", name).Delete(&models.Card{Name: name}).Error
	} else {
		err = m.db.Model(&models.Card{}).Where("name = ?", name).Delete(&models.Card{Name: name}).Error
	}
	if err != nil {
		return fmt.Errorf("%s.DeleteCard: %w", m.Name(), err)
	}
	return nil
}
