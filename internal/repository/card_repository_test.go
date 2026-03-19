package repository

import (
	"store_bot/internal/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getCardRepository(t *testing.T) *CardRepository {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	rep, err := NewCardRepository(&CardRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func TestCreateCard(t *testing.T) {
	rep := getCardRepository(t)

	card, err := rep.CreateCard(&models.Card{Name: "sgamga"})
	assert.Nil(t, err, "unable to create Card")
	assert.NotNil(t, card, "unable to create Card")

	err = rep.DeleteCard("sgamga", true)
	assert.Nil(t, err, "unable to delete Card")
}

func TestGetCard(t *testing.T) {
	rep := getCardRepository(t)
	card, err := rep.CreateCard(&models.Card{Name: "sgamga", ModeratorId: 78})
	assert.Nil(t, err, "unable to create Card")
	assert.NotNil(t, card, "unable to create Card")

	card, err = rep.GetCard(78)
	assert.Nil(t, err, "unable to get Card")
	assert.NotNil(t, card, "unable to get Card")
	err = rep.DeleteCard("sgamga", true)
	assert.Nil(t, err, "unable to delete Card")

	card, err = rep.GetCard(78)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, card, "get un-existen Card")
}

func TestUpdateCard(t *testing.T) {
	rep := getCardRepository(t)

	expectedCard, err := rep.CreateCard(&models.Card{Name: "sgamga", ModeratorId: 78})
	assert.Nil(t, err, "unable to create Card")
	assert.NotNil(t, expectedCard, "unable to create Card")

	expectedCard.Desrciption = "sgamga"
	expectedCard.Name = "sgamga"
	err = rep.UpdateCard(expectedCard)
	assert.Nil(t, err, "unable to udpate Card")

	actualCard, err := rep.GetCard(78)
	assert.Nil(t, err, "unable to get Card after updating")
	assert.Equal(t, expectedCard.Desrciption, actualCard.Desrciption, "card Desrciption isnt match after updating")
	assert.Equal(t, expectedCard.Name, actualCard.Name, "card name isnt match after updating")

	err = rep.DeleteCard("sgamga", true)
	assert.Nil(t, err, "unable to delete Card")
}
