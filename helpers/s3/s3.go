package s3

import (
	"mime/multipart"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// upload file to s3
func UploadFileHelper(file multipart.File, fileName string) (string, error) {
	var cfg = config.GetConfig()

	// The session the S3 Uploader will use
	// s3Config := &aws.Config{
	// 	Region: aws.String(cfg.S3.Region),
	// }

	s3Config := aws.NewConfig()
	s3Config.Region = aws.String(cfg.S3.Region)
	s3Config.CredentialsChainVerboseErrors = aws.Bool(true)

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return "", err
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// uploader := s3manager.NewUploader(sess)

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.S3.Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return upload.Location, nil
}
