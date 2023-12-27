package gateways

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/Hack-Portal/backend/cmd/config"
)

var accessLink string

func TestS3control(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	} else {
		t.Run("testUploadFile", testUploadFile)
		t.Run("testGetPresignedObjectURL", testGetPresignedObjectURL)
		t.Run("testDeleteFile", testDeleteFile)
	}
}

func testUploadFile(t *testing.T) {
	if client == nil {
		t.Fatal("client is nil")
	}
	file, err := os.ReadFile("test.jpg")
	if err != nil {
		t.Error("file open error", err)
	}

	if len(file) == 0 {
		t.Fatal("file is empty")
	}

	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, 1)
	key, err := fs.UploadFile(context.Background(), file, "test.jpg")
	if err != nil {
		t.Error("upload file error", err)
	}

	if key == fmt.Sprintf("%s/%s", hackathonPrefix, "test.jpg") {
		t.Log("upload file success")
	} else {
		t.Error("upload file error")
	}
}

func testGetPresignedObjectURL(t *testing.T) {
	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, 1)
	url, err := fs.GetPresignedObjectURL(context.Background(), "test.jpg")
	if err != nil {
		t.Error("get presigned url error", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("get presigned url error", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("get presigned url error", err)
	}
	t.Log(resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		t.Log("get presigned url success")
	} else {
		t.Error("get presigned url error")
	}
}

func testDeleteFile(t *testing.T) {
	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, 1)
	err := fs.DeleteFile(context.Background(), "test.jpg")
	if err != nil {
		t.Error("delete file error", err)
	}
}
