package gateways

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/usecases/dai"
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

	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, nil, 1)
	key, err := fs.UploadFile(context.Background(), file, "test.jpg")
	if err != nil {
		t.Error("upload file error", err)
	}

	if key == fmt.Sprintf("%s/%s", "hackathon", "test.jpg") {
		t.Log("upload file success")
	} else {
		t.Error("upload file error")
	}
}

func testGetPresignedObjectURL(t *testing.T) {
	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, nil, 1)
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
	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, nil, 1)
	err := fs.DeleteFile(context.Background(), "test.jpg")
	if err != nil {
		t.Error("delete file error", err)
	}
}

func TestParallelGetPresignedObjectURL(t *testing.T) {
	sample := []dai.ParallelGetPresignedObjectURLInput{
		{
			HackathonID: "d7415564-928b-4c82-8c1a-80727a0ad0b9",
			Key:         "hackathon/d7415564-928b-4c82-8c1a-80727a0ad0b9.actions.png",
		},
		{
			HackathonID: "20fc5177-4067-49b3-bed4-8aa3655a1b9b",
			Key:         "hackathon/20fc5177-4067-49b3-bed4-8aa3655a1b9b.actions.png",
		},
		{
			HackathonID: "97ad9bc4-c0ac-46d0-a9a8-44fd21004613",
			Key:         "hackathon/97ad9bc4-c0ac-46d0-a9a8-44fd21004613.actions.png",
		},
		{
			HackathonID: "2d7f714b-02ae-41eb-a72f-9486247bbbe6",
			Key:         "hackathon/2d7f714b-02ae-41eb-a72f-9486247bbbe6.actions.png",
		},
	}

	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client, nil, 1)

	data, err := fs.ParallelGetPresignedObjectURL(context.Background(), sample)
	if err != nil {
		t.Error("parallel get presigned url error", err)
	}

	for i, v := range data {
		t.Log(i, v)
	}
}
