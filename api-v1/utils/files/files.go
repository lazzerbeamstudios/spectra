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

type FileClientStruct struct {
	bucket        string
	project       string
	storageClient *storage.Client
}

var FileClient *FileClientStruct

func (fileClient *FileClientStruct) Upload(fileDir string, fileBlob []byte) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	storageWriter := fileClient.storageClient.Bucket(fileClient.bucket).Object(fileDir).NewWriter(ctx)
	fileBuffer := bytes.NewBuffer(fileBlob)

	_, err := io.Copy(storageWriter, fileBuffer)
	if err != nil {
		return "", err
	}

	err = storageWriter.Close()
	if err != nil {
		return "", err
	}

	url := "https://storage.googleapis.com/" + fileClient.bucket + "/" + fileDir
	return url, nil
}

func SetStorage(credentials string, project string, bucket string) {
	credentialsByte, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		panic("Cannot decode credentials.")
	}

	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsJSON(credentialsByte))
	if err != nil {
		panic("Cannot create storage client.")
	}

	FileClient = &FileClientStruct{
		bucket:        bucket,
		project:       project,
		storageClient: storageClient,
	}
}
