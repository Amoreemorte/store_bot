package repository

import (
	"store_bot/internal/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getModeratorRepository(t *testing.T) *ModeratorRepository {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	rep, err := NewModeratorRepository(&ModeratorRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func TestCreateModerator(t *testing.T) {
	rep := getModeratorRepository(t)
	moderator, err := rep.CreateModerator(&models.Moderator{ID: 1})
	assert.Nil(t, err, "unable to create moderator")
	assert.NotNil(t, moderator, "unable to create moderator")
	err = rep.DeleteModerator(1, true)
	assert.Nil(t, err, "unable to delete moderator")
}

func TestGetModerator(t *testing.T) {
	rep := getModeratorRepository(t)
	moderator, err := rep.CreateModerator(&models.Moderator{ID: 1})
	assert.Nil(t, err, "unable to create moderator")
	assert.NotNil(t, moderator, "unable to create moderator")

	moderator, err = rep.GetModerator(1)
	assert.Nil(t, err, "unable to get moderator")
	assert.NotNil(t, moderator, "unable to create moderator")
	err = rep.DeleteModerator(1, true)
	assert.Nil(t, err, "unable to delete moderator")

	moderator, err = rep.GetModerator(0)
	t.Logf("moderator: %+v, err: %+v", moderator, err)
}
