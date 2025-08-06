package storage

import (
	"fmt"
	"mime/multipart"
	"slices"
	"time"

	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	constants *bootstrap.Constants
	storage   *bootstrap.S3
	clients   *s3.S3
	uploader  *s3manager.Uploader
	buckets   map[enum.BucketType]string
}

func NewS3Storage(
	constants *bootstrap.Constants,
	storage *bootstrap.S3,
) *S3Storage {
	buckets := make(map[enum.BucketType]string)
	buckets[enum.VATTaxpayerCertificate] = storage.Buckets.VATTaxpayerCertificate
	buckets[enum.OfficialNewspaperAD] = storage.Buckets.OfficialNewspaperAD
	buckets[enum.ProfilePic] = storage.Buckets.ProfilePic
	buckets[enum.TicketImage] = storage.Buckets.TicketImage
	buckets[enum.LogoPic] = storage.Buckets.LogoPic
	buckets[enum.NewsMedia] = storage.Buckets.NewsMedia
	return &S3Storage{
		constants: constants,
		storage:   storage,
		buckets:   buckets,
	}
}

func (s3StorageS3Storage *S3Storage) setS3Client(bucketType enum.BucketType) error {
	bucketTypes := enum.GetAllBucketTypes()
	if !slices.Contains(bucketTypes, bucketType) {
		return fmt.Errorf("bucket not exist")
	}
	if s3StorageS3Storage.uploader != nil && s3StorageS3Storage.clients != nil {
		return nil
	}
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s3StorageS3Storage.storage.AccessKey, s3StorageS3Storage.storage.SecretKey, ""),
		Region:      aws.String(s3StorageS3Storage.storage.Region),
		Endpoint:    aws.String(s3StorageS3Storage.storage.Endpoint),
	})

	if err != nil {
		return fmt.Errorf("unable to create AWS session, %w", err)
	}

	s3StorageS3Storage.uploader = s3manager.NewUploader(sess)
	s3StorageS3Storage.clients = s3.New(sess)
	return nil
}

func (s3StorageS3Storage *S3Storage) UploadObject(bucketType enum.BucketType, key string, file *multipart.FileHeader) error {
	err := s3StorageS3Storage.setS3Client(bucketType)
	if err != nil {
		return err
	}
	bucket := s3StorageS3Storage.buckets[bucketType]

	fileReader, err := file.Open()
	if err != nil {
		return fmt.Errorf("unable to open file %q, %w", file.Filename, err)
	}
	defer fileReader.Close()

	_, err = s3StorageS3Storage.clients.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && (aerr.Code() == s3.ErrCodeNoSuchBucket || aerr.Code() == "NotFound") {
			_, err = s3StorageS3Storage.clients.CreateBucket(&s3.CreateBucketInput{
				Bucket: aws.String(bucket),
			})
			if err != nil {
				return fmt.Errorf("unable to create bucket %q, %w", bucket, err)
			}

			err = s3StorageS3Storage.clients.WaitUntilBucketExists(&s3.HeadBucketInput{
				Bucket: aws.String(bucket),
			})
			if err != nil {
				return fmt.Errorf("unable to confirm bucket %q exists, %w", bucket, err)
			}
		} else {
			return fmt.Errorf("unable to check bucket %q, %w", bucket, err)
		}
	}

	_, err = s3StorageS3Storage.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   fileReader,
	})
	if err != nil {
		return fmt.Errorf("unable to upload %q to %q, %w", file.Filename, bucket, err)
	}
	return nil
}

func (s3StorageS3Storage *S3Storage) DeleteObject(bucketType enum.BucketType, key string) error {
	err := s3StorageS3Storage.setS3Client(bucketType)
	if err != nil {
		return err
	}
	bucket := s3StorageS3Storage.buckets[bucketType]

	_, err = s3StorageS3Storage.clients.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to delete %q from %q, %w", key, bucket, err)
	}

	err = s3StorageS3Storage.clients.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to confirm object %q not exists, %w", key, err)
	}
	return nil
}

func (s3StorageS3Storage *S3Storage) GetPresignedURL(bucketType enum.BucketType, objectKey string, expiration time.Duration) (string, error) {
	err := s3StorageS3Storage.setS3Client(bucketType)
	if err != nil {
		return "", err
	}
	bucket := s3StorageS3Storage.buckets[bucketType]

	req, _ := s3StorageS3Storage.clients.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url, nil
}
