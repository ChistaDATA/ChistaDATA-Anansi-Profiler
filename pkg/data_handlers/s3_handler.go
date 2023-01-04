package data_handlers

import (
	"context"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type S3Handler struct {
	s3Config *stucts.S3Config
	s3Client *s3.Client
}

func InitS3Handler(config *stucts.S3Config) *S3Handler {
	return &S3Handler{
		s3Config: config,
		s3Client: createS3Client(config.AccessKeyID, config.SecretAccessKey, config.SessionToken),
	}
}

func createS3Client(key_id string, access_key string, access_token string) *s3.Client {
	// Create an Amazon S3 service client
	return s3.New(s3.Options{
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(key_id, access_key, access_token)),
	})
}

func (h S3Handler) DownloadToFile(s3FileLocation string, downloadFilePath string) error {
	region, bucket, key := splitToRegionBucketKey(s3FileLocation)

	result, err := h.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(options *s3.Options) {
		options.Region = region
	})

	if err != nil {
		log.Error("Couldn't download large object from %v:%v. Here's why: %v\n",
			bucket, key, err)
		return err
	}
	defer result.Body.Close()

	file, err := os.Create(downloadFilePath)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", key, err)
		return err
	}

	_, err = file.Write(body)

	return nil
}

func (h S3Handler) DownloadToFileBig(s3FileLocation string, downloadFilePath string) (*os.File, error) {
	region, bucket, key := splitToRegionBucketKey(s3FileLocation)
	file, err := os.Create(downloadFilePath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer file.Close()

	var partMiBs int64 = 10
	downloader := manager.NewDownloader(h.s3Client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
		d.ClientOptions = append(d.ClientOptions, func(options *s3.Options) {
			options.Region = region
		})
	})

	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Error("Couldn't download large object from %v:%v. Here's why: %v\n",
			bucket, key, err)
		return nil, err
	}

	return file, nil
}

func splitToRegionBucketKey(fullPath string) (string, string, string) {
	//https://profiler-resources.s3.ap-northeast-1.amazonaws.com/s100.log
	s := fullPath[len("https://"):]
	bucketAndRegion := strings.SplitN(s, ".", 5)
	bucket := bucketAndRegion[0]
	region := bucketAndRegion[2]
	d := strings.SplitN(s, "/", 2)
	key := d[1]
	return region, bucket, key
}
