package solr

import "testing"

func TestCollectionCreate(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")
	
	client := NewClient(config)

	response, err := client.Collection.Create(CollectionCreate{
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

	if response.Header.Status != 0 {
		t.Errorf("failed to create collection %v", err)
	}
}

func TestCollectionReload(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Reload(CollectionReload{
		Name:           "collection-test",
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
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Modify(CollectionModifyCollection{
		Collection: "collection-test",
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
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.List()
	if err != nil {
		t.Errorf("failed to list collections %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to list collections %v", err)
	}
}

func TestCollectionProp(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.CollectionProp(CollectionProp{
		Name:           "collection-test",
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
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Rename(CollectionRename{
		Name:           "collection-test",
		Target:         "collection-test",
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
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Create(CollectionCreate{
		Name:                 "collection-test-migrate",
		NumShards:            1,
		ReplicationFactor: 	  1,
		CollectionConfigName: "_default",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection collection-test-migrate %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to create collection collection-test-migrate %v", err)
	}

	response, err = client.Collection.Migrate(CollectionMigrate{
		Collection:       "collection-test",
		TargetCollection: "collection-test-migrate",
		SplitKey:         "a!",
		ForwardTimeout:   100000,
		Async:            false,
	})
	if err != nil {
		t.Errorf("failed to migrate collection %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to migrate collection %v", err)
	}
}

//func TestCollectionBackup(t *testing.T) {
//
//	config := NewConfig().
//		SetUrl("http://127.0.0.1:8981")
//
//	backupFilePath := fmt.Sprintf("bkp_%x", md5.Sum([]byte(time.Now().String())))[0:10]
//
//	client := NewClient(config)
//
//	response, err := client.Collection.Backup(CollectionBackup{
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
//	if response.Header.Status != 0 {
//		t.Errorf("failed to backup collection %v", response)
//	}
//}


//func TestCollectionRestore(t *testing.T) {
//	config := NewConfig().
//		SetUrl("http://127.0.0.1:8981")
//
//	backupFilePath := fmt.Sprintf("bkp_%x", md5.Sum([]byte(time.Now().String())))[0:8]
//
//	client := NewClient(config)
//
//	response, err := client.Collection.Backup(CollectionBackup{
//		Collection:     "collection-test",
//		Name:           backupFilePath,
//		Location:       "/tmp/",
//		Async:          false,
//	})
//	if err != nil {
//		t.Errorf("failed to backup collection %v", err)
//	}
//
//	if response.Header.Status != 0 {
//		t.Errorf("failed to backup collection %v", response)
//	}
//
//	response, err = client.Collection.Restore(CollectionRestore{
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
//	if response.Header.Status != 0 {
//		t.Errorf("failed to restore collection %v", response)
//	}
//}

func TestCollectionDelete(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Delete(CollectionDelete{
		Name:           "collection-test",
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
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	response, err := client.Collection.Delete(CollectionDelete{
		Name:           "collection-test-migrate",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection migrate %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to delete collection migrate %v", err)
	}
}

