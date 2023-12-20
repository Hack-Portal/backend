package gateways

import (
	"bytes"
	"context"
	"io"
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
		t.Run("testDeleteFile", testDeleteFile)
	}
}

func testUploadFile(t *testing.T) {
	file, err := os.ReadFile("test.jpg")
	if err != nil {
		t.Error("file open error", err)
	}

	if len(file) == 0 {
		t.Fatal("file is empty")
	}

	t.Log("file size:", len(file))

	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client)
	accessLink, err = fs.UploadFile(context.Background(), file, "test.jpg")
	if err != nil {
		t.Error("upload file error", err)
	}
	t.Log(accessLink)

	// check file

	req, err := http.NewRequest("GET", accessLink, nil)
	if err != nil {
		t.Error("request error", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("request error", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("request error", err)
	}

	// テスト通るのかわからん

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error("read error", err)
	}

	if len(data) == 0 {
		t.Error("data is empty")
	}

	if bytes.Compare(data, file) != 0 {
		t.Error("data is not same")
	}
}

func testDeleteFile(t *testing.T) {
	fs := NewCloudflareR2(config.Config.Buckets.Bucket, client)
	err := fs.DeleteFile(context.Background(), "test.jpg")
	if err != nil {
		t.Error("delete file error", err)
	}
}
