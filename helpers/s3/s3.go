package s3

import (
	"mime/multipart"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// upload file to s3
func UploadFileHelper(file multipart.File, fileName string) (string, error) {
	var cfg = config.GetConfig()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AwsS3.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AwsS3.AwsId,
			cfg.AwsS3.AwsKey,
			"",
		),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.AwsS3.Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return upload.Location, nil
}
