package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	CardId int64     `gorm:"not null"`
	Card   *Card     `gorm:"not null;foreignKey:CardId;references:ID"`
	gorm.Model
}

type Collection struct {
	Name  string  `gorm:"primaryKey"`
	Cards []*Card `gorm:"not null;foreignKey:CollectionId;references:ID"`

	gorm.Model
}

type Card struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"size:50"`
	Desrciption string `gorm:"size:2048"`
	Price       int

	CollectionId int64
	Collection   *Collection `gorm:"not null;foreignKey:CollectionId;references:ID"`

	Photos []*Photo `gorm:"not null;foreignKey:CardId;references:ID"`
	Orders []*Order `gorm:"not null;foreignKey:CardId;references:ID"`

	gorm.Model
}

type Order struct {
	ID int64 `gorm:"primaryKey"`

	UserId int64
	User   *User `gorm:"not null;foreignKey:UserId;references:ID"`

	CardId int64
	Card   *Card `gorm:"not null;foreignKey:CardId;references:ID"`

	BucketId int64
	Bucket   *Bucket `gorm:"not null;foreignKey:BucketId;references:ID"`

	IsEnded bool

	gorm.Model
}

type Bucket struct {
	ID int64 `gorm:"primaryKey"`

	OwnerId int64
	User    *User `gorm:"not null;foreignKey:OwnerId;references:ID"`

	Orders []*Order `gorm:"not null;foreignKey:BucketId;references:ID"`

	gorm.Model
}
