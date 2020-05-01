package solr

import (
	"context"
	"testing"
)

func TestCollectionCreate(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Create(context.Background(), CollectionCreate{
		Name:                 "collection-test",
		RouterName:           "compositeId",
		NumShards:            1,
		ReplicationFactor: 	  1,
		CollectionConfigName: "_default",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create collection %v", err)
	}
}

func TestCollectionReload(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Reload(context.Background(), CollectionReload{
		Name:           "collection-test",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to reload collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to reload collection %v", err)
	}
}

func TestCollectionModify(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Modify(context.Background(), CollectionModifyCollection{
		Collection: "collection-test",
		MaxShardsPerNode:      1,
	})
	if err != nil {
		t.Errorf("failed to modify collections %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to modify collections %v", err)
	}
}

func TestCollectionList(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.List(context.Background())
	if err != nil {
		t.Errorf("failed to list collections %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to list collections %v", err)
	}
}

func TestCollectionProp(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.CollectionProp(context.Background(), CollectionProp{
		Name:           "collection-test",
		PropertyName:   "timestamp",
		PropertyValue:  "dateTime",
	})
	if err != nil {
		t.Errorf("failed to modify property %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to modify property %v", err)
	}
}

func TestCollectionRename(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Rename(context.Background(), CollectionRename{
		Name:           "collection-test",
		Target:         "collection-test",
	})
	if err != nil {
		t.Errorf("failed to rename collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to rename collection %v", err)
	}
}

func TestCollectionMigrate(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Create(context.Background(), CollectionCreate{
		Name:                 "collection-test-migrate",
		NumShards:            1,
		ReplicationFactor: 	  1,
		CollectionConfigName: "_default",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection collection-test-migrate %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create collection collection-test-migrate %v", err)
	}

	response, err = client.Collection.Migrate(context.Background(), CollectionMigrate{
		Collection:       "collection-test",
		TargetCollection: "collection-test-migrate",
		SplitKey:         "a!",
		ForwardTimeout:   100000,
		Async:            false,
	})
	if err != nil {
		t.Errorf("failed to migrate collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to migrate collection %v", err)
	}
}

//func TestCollectionBackup(t *testing.T) {
//
//	client := NewClient()
//
//	backupFilePath := fmt.Sprintf("bkp_%x", md5.Sum([]byte(time.Now().String())))[0:10]
//
//	client := NewClient(config)
//
//	response, err := client.Collection.Backup(context.Background(), CollectionBackup{
//		Collection:     "collection-test",
//		Name:           backupFilePath,
//		Location:       "/tmp",
//		Async:          false,
//	})
//
//	if err != nil {
//		t.Errorf("failed to backup collection %v", err)
//	}
//
//	if response.ResponseHeader.Status != 0 {
//		t.Errorf("failed to backup collection %v", response)
//	}
//}


//func TestCollectionRestore(t *testing.T) {
//	client := NewClient()
//
//	backupFilePath := fmt.Sprintf("bkp_%x", md5.Sum([]byte(time.Now().String())))[0:8]
//
//	client := NewClient(config)
//
//	response, err := client.Collection.Backup(context.Background(), CollectionBackup{
//		Collection:     "collection-test",
//		Name:           backupFilePath,
//		Location:       "/tmp/",
//		Async:          false,
//	})
//	if err != nil {
//		t.Errorf("failed to backup collection %v", err)
//	}
//
//	if response.ResponseHeader.Status != 0 {
//		t.Errorf("failed to backup collection %v", response)
//	}
//
//	response, err = client.Collection.Restore(context.Background(), CollectionRestore{
//		Collection:           backupFilePath,
//		Name:                 backupFilePath,
//		Location:       	  "/tmp/",
//		Async:                false,
//		ReplicationFactor:    1,
//	})
//	if err != nil {
//		t.Errorf("failed to restore collection %v", err)
//	}
//
//	if response.ResponseHeader.Status != 0 {
//		t.Errorf("failed to restore collection %v", response)
//	}
//}

func TestCollectionDelete(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Delete(context.Background(), CollectionDelete{
		Name:           "collection-test",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to delete collection %v", err)
	}
}

func TestCollectionDeleteMigrate(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Delete(context.Background(), CollectionDelete{
		Name:           "collection-test-migrate",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection migrate %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to delete collection migrate %v", err)
	}
}

