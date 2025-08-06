package s3

import (
	"mime/multipart"
	"time"

	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
)

type S3Storage interface {
	DeleteObject(bucketType enum.BucketType, key string) error
	GetPresignedURL(bucketType enum.BucketType, objectKey string, expiration time.Duration) (string, error)
	UploadObject(bucketType enum.BucketType, key string, file *multipart.FileHeader) error
}
