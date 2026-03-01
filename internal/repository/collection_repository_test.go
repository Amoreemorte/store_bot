package repository

import (
	"fmt"
	"store_bot/internal/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getCollectionRepository(t *testing.T) *CollectionRepository {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	rep, err := NewCollectionRepository(&CollectionRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func TestCreateCollection(t *testing.T) {
	rep := getCollectionRepository(t)
	Collection, err := rep.CreateCollection(&models.Collection{Name: "sgamga"})
	assert.Nil(t, err, "unable to create Collection")
	assert.NotNil(t, Collection, "unable to create Collection")
	err = rep.DeleteCollection("sgamga", true)
	assert.Nil(t, err, "unable to delete Collection")
}

func TestGetCollection(t *testing.T) {
	rep := getCollectionRepository(t)
	Collection, err := rep.CreateCollection(&models.Collection{Name: "sgamga"})
	assert.Nil(t, err, "unable to create Collection")
	assert.NotNil(t, Collection, "unable to create Collection")

	Collection, err = rep.GetCollection("sgamga")
	assert.Nil(t, err, "unable to get Collection")
	assert.NotNil(t, Collection, "unable to get Collection")
	err = rep.DeleteCollection("sgamga", true)
	assert.Nil(t, err, "unable to delete Collection")

	Collection, err = rep.GetCollection("sgamga")
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, Collection, "get un-existen Collection")
}

func TestGetCollections(t *testing.T) {
	rep := getCollectionRepository(t)
	CollectionsNames := []string{"sgamga", "door", "sguga"}

	for _, name := range CollectionsNames {
		Collection, err := rep.CreateCollection(&models.Collection{Name: name})
		assert.Nil(t, err, "unable to create Collection")
		assert.NotNil(t, Collection, "unable to create Collection")
	}

	Collections, err := rep.GetCollections()
	assert.Nil(t, err, "unable to get Collections")
	assert.NotNil(t, Collections, "unable to get Collections")
	assert.Equal(t, len(CollectionsNames), len(Collections), fmt.Sprintf("collections: %+v", Collections))

	for _, name := range CollectionsNames {
		err = rep.DeleteCollection(name, true)
		assert.Nil(t, err, "unable to delete Collection")
	}
}
