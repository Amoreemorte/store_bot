package models

import (
	"gorm.io/gorm"
)

// Main entity of moderator bot.
//
// Can create collections, cards etc.
type Moderator struct {
	ID int64 `gorm:"primaryKey"`

	State *ModeratorState `gorm:"not null;foreignKey:ModeratorId;references:ID"`
	gorm.Model
}

type ModeratorStateRef = string

const NoAction ModeratorStateRef = "NoAction"
const Name ModeratorStateRef = "Name"
const Decsription ModeratorStateRef = "Description"
const Price ModeratorStateRef = "Price"
const AddPhoto ModeratorStateRef = "AddPhoto"

type ModeratorState struct {
	State       ModeratorStateRef `gorm:"column: state"`
	ModeratorId int64             `gorm:"primaryKey"`

	Moderator *Moderator `gorm:"not null;foreignKey:ModeratorId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
