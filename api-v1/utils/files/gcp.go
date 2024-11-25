package files

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type StorageClientGoogleStruct struct {
	bucket        string
	project       string
	storageClient *storage.Client
}

var StorageClientGoogle *StorageClientGoogleStruct

func (client *StorageClientGoogleStruct) UploadFile(path string, content []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	writer := client.storageClient.Bucket(client.bucket).Object(path).NewWriter(ctx)
	reader := bytes.NewBuffer(content)

	if _, err := io.Copy(writer, reader); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	url := "https://storage.googleapis.com/" + client.bucket + "/" + path
	return url, nil
}

func SetStorage(credentials string, project string, bucket string) {
	credentialsByte, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		panic("Failed to decode GCP credentials")
	}

	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsJSON(credentialsByte),
	)
	if err != nil {
		panic("Failed to create GCP storage client")
	}

	StorageClientGoogle = &StorageClientGoogleStruct{
		bucket:        bucket,
		project:       project,
		storageClient: storageClient,
	}
}
