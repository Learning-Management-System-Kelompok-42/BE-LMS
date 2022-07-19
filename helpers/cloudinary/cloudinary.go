package cloudOD

import (
	"context"
	"time"

	"github.com/Learning-Management-System-Kelompok-42/BE-LMS/config"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	var cfg = config.GetConfig()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(cfg.Storage.CloudName, cfg.Storage.APIKey, cfg.Storage.APISecret)
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.GetConfig().Storage.UploadStorage})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
