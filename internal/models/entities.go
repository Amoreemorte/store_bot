package models

import (
	"gorm.io/gorm"
)

// Main entity of moderator bot.
//
// Can create collections, cards etc.
type Moderator struct {
	ID    int64 `gorm:"primaryKey"`
	State string

	gorm.Model
}

// Main entity of client bot.
//
// Can create buckets, search cards etc.
type User struct {
	ID int64 `gorm:"primaryKey"`
	// User may not have an username
	Username *string `gorm:"unique"`

	Orders []*Order `gorm:"not null;foreignKey:UserId;references:ID"`
	Bucket *Bucket  `gorm:"not null;foreignKey:OwnerId;references:ID"`

	gorm.Model
}
