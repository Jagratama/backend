package aws

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Uploader(bucket string) (*S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Uploader{
		Client:     client,
		BucketName: bucket,
	}, nil
}

func (u *S3Uploader) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	key := fmt.Sprintf("%s/%d%s", folder, time.Now().UnixNano(), ext)

	_, err := u.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.BucketName),
		Key:         aws.String(key),
		Body:        file,
		ACL:         "public-read",
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.BucketName, key)
	return url, nil
}

func (u *S3Uploader) UploadBuffer(ctx context.Context, buf []byte, filename, folder, contentType string) (string, error) {
	key := fmt.Sprintf("%s/%s", folder, filename)

	_, err := u.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.BucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buf),
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.BucketName, key)
	return url, nil
}
