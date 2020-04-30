package solr

import (
	"testing"
)


func TestCollectionCreate(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Create(CollectionCreate{
		Name:                 "client-solr",
		RouterName:           "compositeId",
		NumShards:            1,
		CollectionConfigName: "client-solr.AUTOCREATED",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to create collection %v", err)
	}
}

func TestCollectionReload(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Reload(CollectionReload{
		Name:           "client-solr",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to reload collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to reload collection %v", err)
	}
}

func TestCollectionModify(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Modify(CollectionModifyCollection{
		Collection: "client-solr",
		MaxShardsPerNode:      1,
	})
	if err != nil {
		t.Errorf("failed to modify collections %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to modify collections %v", err)
	}
}

func TestCollectionList(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.List()
	if err != nil {
		t.Errorf("failed to list collections %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to list collections %v", err)
	}
}

func TestCollectionProp(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.CollectionProp(CollectionProp{
		Name:           "client-solr",
		PropertyName:   "timestamp",
		PropertyValue:  "dateTime",
	})
	if err != nil {
		t.Errorf("failed to modify property %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to modify property %v", err)
	}
}

func TestCollectionRename(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Rename(CollectionRename{
		Name:           "client-solr",
		Target:         "client-solr",
	})
	if err != nil {
		t.Errorf("failed to rename collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to rename collection %v", err)
	}
}

func TestCollectionMigrate(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Create(CollectionCreate{
		Name:                 "client-solr-migrate",
		RouterName:           "compositeId",
		NumShards:            2,
		CollectionConfigName: "_default",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection client-solr-migrate %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to create collection client-solr-migrate %v", err)
	}

	response, err = collection.Migrate(CollectionMigrate{
		Collection:       "client-solr",
		TargetCollection: "client-solr-migrate",
		SplitKey:         "a!",
		ForwardTimeout:   10000,
		Async:            false,
	})
	if err != nil {
		t.Errorf("failed to migrate collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to migrate collection %v", err)
	}
}

func TestCollectionBackup(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Backup(CollectionBackup{
		Collection:     "client-solr",
		Name:           "bkp_3",
		Location:       "/tmp",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to backup collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to backup collection %v", err)
	}
}

func TestCollectionRestore(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Restore(CollectionRestore{
		Collection:           "client-solr-restore",
		Name:                 "bkp_3",
		Location:             "/tmp",
		Async:                false,
		ReplicationFactor:    2,
	})
	if err != nil {
		t.Errorf("failed to restore collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to restore collection %v", err)
	}
}

func TestCollectionDelete(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Delete(CollectionDelete{
		Name:           "client-solr",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to delete collection %v", err)
	}
}

func TestCollectionDeleteMigrate(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981").
		SetPrefix("/solr")

	collection := NewCollection(config)

	response, err := collection.Delete(CollectionDelete{
		Name:           "client-solr-migrate",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection migrate %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to delete collection migrate %v", err)
	}
}

