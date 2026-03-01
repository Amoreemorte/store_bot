package repository

import (
	"fmt"
	"store_bot/internal/models"

	"gorm.io/gorm"
)

type CollectionRepository struct {
	cfg *CollectionRepositoryConfig
	db  *gorm.DB
}

func NewCollectionRepository(cfg *CollectionRepositoryConfig, db *gorm.DB) (*CollectionRepository, error) {
	if cfg.MigrateEnable {
		Collection := &models.Collection{}
		if err := db.AutoMigrate(Collection); err != nil {
			return nil, fmt.Errorf("Unable to migrate collection: %w", err)
		}
	}
	return &CollectionRepository{
		db:  db,
		cfg: cfg,
	}, nil
}

func (c *CollectionRepository) Name() string {
	return "CollectionRepository"
}

func (c *CollectionRepository) CreateCollection(Collection *models.Collection) (*models.Collection, error) {
	if err := c.db.Model(&models.Collection{}).Create(Collection).Error; err != nil {
		return nil, fmt.Errorf("%s.CreateCollection: %w", c.Name(), err)
	}
	return Collection, nil
}

func (c *CollectionRepository) GetCollection(name string) (*models.Collection, error) {
	Collection := &models.Collection{Name: name}
	if err := c.db.Model(&models.Collection{}).Where(Collection).First(Collection).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetCollection: %w", c.Name(), err)
	}
	return Collection, nil
}

func (c *CollectionRepository) GetCollections() ([]models.Collection, error) {
	collections := make([]models.Collection, 0)
	if err := c.db.Model(&models.Collection{}).Find(&collections).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s.GetCollections: %w", c.Name(), err)
	}
	return collections, nil
}

// deleteUnscoped = false will be used soft deleting
//
// deleteUnscoped = true will be used full deleting
func (m *CollectionRepository) DeleteCollection(name string, deleteUnscoped bool) error {
	var err error
	if deleteUnscoped {
		err = m.db.Unscoped().Where("name = ?", name).Delete(&models.Collection{}).Error
	} else {
		err = m.db.Where("name = ?", name).Delete(&models.Collection{}).Error
	}
	if err != nil {
		return fmt.Errorf("%s.DeleteCollection: %w", m.Name(), err)
	}
	return nil
}
