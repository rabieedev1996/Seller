package Service

import (
	"Seller/Seller.Application/Model"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"io/fs"
	"io/ioutil"
	"strings"
)

type ArvanCloudUploadFileService struct {
}

func (service ArvanCloudUploadFileService) UploadFile(model Model.UploadFileServiceInputDTO) (Model.UploadFileServiceOutputDTO, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("842303f9-f239-46c5-8876-19f6c6b038ac", "1176fa34750770f0546f7325eb3200871af93c6941f1fe2d02d1e65ba07962ba", ""),
		Endpoint:    aws.String("https://s3.ir-thr-at1.arvanstorage.com"),
	}))
	model.Path = strings.Replace(strings.TrimRight(strings.TrimRight(model.Path, "\\"), "/"), "\\", "/", -1)

	guid := uuid.New().String()
	fileName := model.Path + "/" + guid + "." + model.Extension
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	reader := bytes.NewReader(model.FileBytes)

	// Upload the file to S3.
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("goseller"),
		Key:    aws.String(fileName[1:]),
		Body:   reader,
		ACL:    aws.String("public-read"),
	})
	output := Model.UploadFileServiceOutputDTO{
		RelativeUrl: fileName,
		BaseUrl:     "https://s3.ir-thr-at1.arvanstorage.com/goseller",
	}
	return output, err
}

type DiskStorageUploadFileService struct {
}

func (service DiskStorageUploadFileService) UploadFile(model Model.UploadFileServiceInputDTO) (Model.UploadFileServiceOutputDTO, error) {
	basePath := ""
	guid := uuid.New().String()
	fileName := basePath + "\\" + model.Path + "\\" + guid + "." + model.Extension
	// Create an uploader with the session and default options

	permissions := 0644 // or whatever you need
	err := ioutil.WriteFile(fileName, model.FileBytes, fs.FileMode(permissions))

	output := Model.UploadFileServiceOutputDTO{
		RelativeUrl: fileName,
		BaseUrl:     "",
	}
	return output, err
}
