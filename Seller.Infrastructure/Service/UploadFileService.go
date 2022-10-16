package Service

import (
	"Seller/Seller.Application/Model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"io"
	"io/fs"
	"io/ioutil"
	"strings"
)

type ArvanCloudUploadFileService struct {
}

func (service ArvanCloudUploadFileService) UploadFile(model Model.UploadFileServiceInputDTO) (Model.UploadFileServiceOutputDTO, error) {
	sess := session.Must(session.NewSession())
	model.Path = strings.Replace(strings.TrimRight(strings.TrimRight(model.Path, "\\"), "/"), "\\", "/", -1)

	guid := uuid.New().String()
	fileName := model.Path + "/" + guid + "." + model.Extension
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	var reader io.Reader
	reader.Read(model.FileBytes)

	// Upload the file to S3.
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("GoSeller"),
		Key:    aws.String(fileName),
		Body:   reader,
	})
	output := Model.UploadFileServiceOutputDTO{
		RelativeUrl: fileName,
		BaseUrl:     "",
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
