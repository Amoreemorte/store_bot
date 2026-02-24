package models

import "time"

// Main entity of moderator bot.
//
// Can create collections, cards etc.
type Moderator struct {
	ID    int64
	State string

	CreatedAt      time.Duration
	LastSeenOnline time.Duration
}

// Main entiry of client bot.
//
// Can create buckets, search cards etc.
type User struct {
	ID int64
	// User may not have an username
	Username *string
}
