package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID
	CreatedAt time.Duration
}

type Collection struct {
	Name string

	CreatedAt time.Duration
}

type Card struct {
	Collection  *Collection
	Name        string
	Desrciption string
	Price       int
	Photos      []*Photo

	CreatedAt time.Duration
}

type Order struct {
	ID          int64
	RelatedCard *Card
	OrderFrom   *User

	CreatedAt time.Duration
}

type Busket struct {
	ID     int64
	Orders []*Order
	Owner  *User
}
