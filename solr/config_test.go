package solr

import (
	"context"
	"testing"
)

func TestUploadConfig(t *testing.T) {
	client := NewClient()
	_, err := client.Config.Upload(context.Background(), "../example/configs/identify-events.zip", "identify-events")
	if err != nil {
		t.Errorf("failed to upload config %v", err)
	}
}

func TestCreateConfig(t *testing.T) {
	client := NewClient()
	_, err := client.Config.Create(context.Background(), CreateConfig{
		Create: Config{
			Name:                   "identify-events.CREATE",
			BaseConfigSet:          "identify-events",
		},
	})
	if err != nil {
		t.Errorf("failed to create config %v", err)
	}
}

func TestDeleteConfig1(t *testing.T) {
	client := NewClient()
	_, err := client.Config.Delete(context.Background(), "identify-events.CREATE")
	if err != nil {
		t.Errorf("failed to delete config %v", err)
	}
}

func TestDeleteConfig2(t *testing.T) {
	client := NewClient()
	_, err := client.Config.Delete(context.Background(), "identify-events")
	if err != nil {
		t.Errorf("failed to delete config %v", err)
	}
}
