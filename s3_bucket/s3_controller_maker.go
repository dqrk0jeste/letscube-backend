package s3_bucket

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dqrk0jeste/letscube-backend/util"
)

type S3Controller struct {
	uploader *manager.Uploader
	client   *s3.Client
}

func ControllerMaker() (*S3Controller, error) {
	awsConfig, err := aws_config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsConfig)
	uploader := manager.NewUploader(client)

	controller := &S3Controller{
		uploader: uploader,
		client:   client,
	}

	return controller, nil
}

func (controller *S3Controller) Upload(context context.Context, file *multipart.FileHeader, nameToSaveAs string) (*manager.UploadOutput, error) {
	jpegImage, err := util.ConvertToJPEG(file)
	if err != nil {
		return nil, err
	}

	uploadedFile, err := controller.uploader.Upload(context, &s3.PutObjectInput{
		Bucket: aws.String("letscube"),
		Key:    aws.String(nameToSaveAs),
		Body:   jpegImage,
	})
	if err != nil {
		return nil, err
	}

	return uploadedFile, nil
}

func (controller *S3Controller) Delete(context context.Context, nameOfTheFile string) error {
	_, err := controller.client.DeleteObject(context, &s3.DeleteObjectInput{
		Bucket: aws.String("letscube"),
		Key:    aws.String(nameOfTheFile),
	})
	if err != nil {
		return err
	}

	return nil
}
