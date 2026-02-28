package repository

import (
	"store_bot/internal/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getModeratorStateRepository(t *testing.T) *ModeratorStateRepository {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.Exec("PRAGMA foreign_keys = ON;")
	rep, err := NewModeratorStateRepository(&ModeratorStateRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func TestCreateModeratorState(t *testing.T) {
	rep := getModeratorStateRepository(t)
	mdRep := getModeratorRepository(t)
	t.Run("Create moderator state for non-existen moderator", func(t *testing.T) {
		state, err := rep.CreateModeratorState(&models.ModeratorState{ModeratorId: 78, State: models.NoAction})
		assert.Nil(t, state)
		assert.NotNil(t, err)
	})
	t.Run("Create moderator state for existen moderator", func(t *testing.T) {
		_, err := mdRep.CreateModerator(&models.Moderator{ID: 78})
		assert.Nil(t, err)

		state, err := rep.CreateModeratorState(&models.ModeratorState{ModeratorId: 78, State: models.NoAction})
		assert.NotNil(t, state)
		assert.Nil(t, err)

		err = rep.DeleteModeratorState(78, true)
		assert.Nil(t, err)
		err = mdRep.DeleteModerator(78, true)
		assert.Nil(t, err)
	})
}

func TestGetModeratorState(t *testing.T) {
	rep := getModeratorStateRepository(t)
	mdRep := getModeratorRepository(t)
	t.Run("Get moderator state for non-existen moderator", func(t *testing.T) {
		state, err := rep.GetModeratorState(78)
		assert.Nil(t, state)
		assert.NotNil(t, err)
	})
	t.Run("Create moderator state for existen moderator", func(t *testing.T) {
		_, err := mdRep.CreateModerator(&models.Moderator{ID: 78})
		assert.Nil(t, err)
		state, err := rep.CreateModeratorState(&models.ModeratorState{ModeratorId: 78, State: models.NoAction})
		assert.NotNil(t, state)
		assert.Nil(t, err)

		state, err = rep.GetModeratorState(78)
		assert.NotNil(t, state)
		assert.Nil(t, err)

		err = rep.DeleteModeratorState(78, true)
		assert.Nil(t, err)
		err = mdRep.DeleteModerator(78, true)
		assert.Nil(t, err)
	})
}
