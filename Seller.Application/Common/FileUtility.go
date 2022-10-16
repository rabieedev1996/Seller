package Common

import (
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Model"
	"Seller/Seller.Domain/Entities"
	gomime "github.com/cubewise-code/go-mime"
	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	ptime "github.com/yaa110/go-persian-calendar"
	"strconv"
	"time"
)

const (
	OTHER_FILE = iota
	PRODUCT_FILE
)

type FileUtility struct {
	UploadFileService Infrastructure.IUploadFileService
	IFileRepository   Presistence.IFileRepository
}

func (utility FileUtility) GetRelPath(relationType int) string {
	switch relationType {
	case OTHER_FILE:
		return "\\OtherFile"
	case PRODUCT_FILE:
		return "\\ProductFile"
	default:
		return "\\UnknownFile"
	}
}

func (utility FileUtility) UploadFile(file Model.UploadFileUtilityInputDTO) Entities.File {
	uploadFileServiceModel := Model.UploadFileServiceInputDTO{}
	mapper.AutoMapper(&file, &uploadFileServiceModel)
	uploadFileServiceModel.Path = utility.GetRelPath(file.RelationType)
	uploadResult, _ := utility.UploadFileService.UploadFile(uploadFileServiceModel)
	jalaliDate := ptime.Now()
	numericJalaliDate, _ := strconv.Atoi(jalaliDate.Format("yyyyMMdd"))
	fileEntity := Entities.File{
		Id:                uuid.New().String(),
		File_Name:         file.FileName,
		File_Length:       file.FileLength,
		Content_Type:      gomime.TypeByExtension(file.Extension),
		Extension:         file.Extension,
		Relation_Type:     file.RelationType,
		Relation_Id:       file.RelationId,
		Created_At:        time.Now(),
		Jalali_Created_At: numericJalaliDate,
		Is_Deleted:        false,
		Relative_Url:      uploadResult.RelativeUrl,
		Base_Url:          uploadResult.RelativeUrl,
	}
	utility.IFileRepository.Create(&fileEntity)
	return fileEntity
}
