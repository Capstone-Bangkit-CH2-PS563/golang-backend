package util

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gosimple/slug"
)

const (
	projectID  = "nutriplant"  // FILL IN WITH YOURS
	bucketName = "nutriplant" // FILL IN WITH YOURS
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func UploadFile(file multipart.File, object string) (string, error) {
	os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "img/",
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	newObject := slug.Make(object) + filepath.Ext(object)
	// Upload an object with storage.Writer.
	wc := uploader.cl.Bucket(uploader.bucketName).Object(uploader.uploadPath + newObject).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	fileUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s%s", uploader.bucketName, uploader.uploadPath, newObject)
	return fileUrl, nil
}


func UploadFileToScan(file multipart.File, object string) (string, error) {
	os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "scan/",
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	newObject := slug.Make(object) + filepath.Ext(object)
	// Upload an object with storage.Writer.
	wc := uploader.cl.Bucket(uploader.bucketName).Object(uploader.uploadPath + newObject).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	fileUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s%s", uploader.bucketName, uploader.uploadPath, newObject)
	return fileUrl, nil
}