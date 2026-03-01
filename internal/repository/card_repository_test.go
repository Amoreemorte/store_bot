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
	assert.Nil(t, err, "unable to create card")
	assert.NotNil(t, card, "unable to create card")
	err = rep.DeleteCard("sgamga", true)
	assert.Nil(t, err, "unable to delete card")
}

func TestGetCard(t *testing.T) {
	rep := getCardRepository(t)
	card, err := rep.CreateCard(&models.Card{Name: "sgamga"})
	assert.Nil(t, err, "unable to create card")
	assert.NotNil(t, card, "unable to create card")

	card, err = rep.GetCard("sgamga")
	assert.Nil(t, err, "unable to get card")
	assert.NotNil(t, card, "unable to get card")
	err = rep.DeleteCard("sgamga", true)
	assert.Nil(t, err, "unable to delete card")

	card, err = rep.GetCard("sgamga")
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, card, "get un-existen card")
}

func TestGetCards(t *testing.T) {
	rep := getCardRepository(t)
	cardsNames := []string{"sgamga", "door", "sguga"}

	for _, name := range cardsNames {
		card, err := rep.CreateCard(&models.Card{Name: name})
		assert.Nil(t, err, "unable to create card")
		assert.NotNil(t, card, "unable to create card")
	}

	cards, err := rep.GetCards()
	assert.Nil(t, err, "unable to get cards")
	assert.NotNil(t, cards, "unable to get cards")
	assert.Equal(t, len(cardsNames), len(cards))

	for _, name := range cardsNames {
		err = rep.DeleteCard(name, true)
		assert.Nil(t, err, "unable to delete card")
	}
}
