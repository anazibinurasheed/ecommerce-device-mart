package helper

import (
	"bytes"
	"fmt"
	"log"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// Todo: validation of file that are going to upload.
func UploadMediaToS3(data []byte, filename string) (string, error) {
	// creating unique file name
	filename = uuid.New().String() + filename
	key := config.GetConfig().S3BucketMediaPath + filename
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.GetConfig().AWSRegion),

		Credentials: credentials.NewStaticCredentials(
			config.GetConfig().AWSAccessKeyID, config.GetConfig().AWSSecretAccessKey, ""),
		Endpoint: nil,
	},
	)

	if err != nil {
		return "", err
	}

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	uploader := s3manager.NewUploader(sess)

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.GetConfig().S3BucketName),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(key),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: bytes.NewReader(data),
	})

	if err != nil {
		// Print the error and exit.
		log.Printf("Unable to upload %q to %q, %v", filename, config.GetConfig().S3BucketName, err)
		return "", err
	}

	url := fmt.Sprintf("https://s3.amazonaws.com/%s/%s", config.GetConfig().S3BucketName, key)
	fmt.Println(url)

	return output.Location, nil

}
