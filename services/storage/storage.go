package storage

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func getClient() (*minio.Client, error) {
	spacesKey := os.Getenv("SPACES_KEY")
	spacesSecret := os.Getenv("SPACES_SECRET")
	endpoint := os.Getenv("SPACES_ENDPOINT")
	// if there is a colon, assume its not ssl, ie localhost:9000
	useSSL := !strings.Contains(endpoint, ":")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(spacesKey, spacesSecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func MakeBucket(ctx context.Context, bucketName string) error {
	minioClient, err := getClient()
	if err != nil {
		return err
	}

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		return err
	}

	return nil
}

func Upload(ctx context.Context, bucketName string, objectName string, body io.Reader, size int64, contentType string) error {
	minioClient, err := getClient()
	if err != nil {
		return err
	}

	_, err = minioClient.PutObject(ctx, bucketName, objectName, body, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func UploadBytes(ctx context.Context, bucketName string, objectName string, contentBytes []byte) error {
	buffer := bytes.NewBuffer(contentBytes)

	return Upload(ctx, bucketName, objectName, buffer, int64(len(buffer.Bytes())), "application/octet-stream")
}

func Download(ctx context.Context, bucketName string, objectName string) (*minio.Object, error) {
	minioClient, err := getClient()
	if err != nil {
		return nil, err
	}

	obj, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	return obj, err
}

func DownloadBytes(ctx context.Context, bucketName string, objectName string) ([]byte, error) {
	obj, err := Download(ctx, bucketName, objectName)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(obj)
}
