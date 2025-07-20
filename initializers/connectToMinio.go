package initializers

import (
	"context"
	"log"

	"webimg/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func ConnectToMinio() {
	ctx := context.Background()
	cfg := config.GetConfig()

	// Initialize minio client object.
	var err error
	MinioClient, err = minio.New(cfg.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioConfig.AccessKeyID, cfg.MinioConfig.SecretAccessKey, ""),
		Secure: cfg.MinioConfig.UseSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	err = MinioClient.MakeBucket(ctx, cfg.MinioConfig.BucketName, minio.MakeBucketOptions{Region: cfg.MinioConfig.Location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinioClient.BucketExists(ctx, cfg.MinioConfig.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", cfg.MinioConfig.BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", cfg.MinioConfig.BucketName)
	}
}
