package storage

import (
	"bytes"
	"context"
	"testing"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	bucketName := "bucket1234"
	objectName := "object5678"
	contentBytes := []byte("abc def")

	err := MakeBucket(ctx, bucketName)
	if err != nil {
		t.Fatalf("error making bucket: %s", err)
	}

	err = UploadBytes(ctx, bucketName, objectName, contentBytes)
	if err != nil {
		t.Fatalf("error uploading: %s", err)
	}

	downloadedBytes, err := DownloadBytes(ctx, bucketName, objectName)
	if err != nil {
		t.Fatalf("error downloading: %s", err)
	}

	if bytes.Compare(contentBytes, downloadedBytes) != 0 {
		t.Fatalf("downloaded does not match original")
	}
}
