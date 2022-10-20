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
	"strings"
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

func (utility FileUtility) UploadFile(file Model.UploadFileUtilityInputDTO) *Entities.File {
	uploadFileServiceModel := Model.UploadFileServiceInputDTO{}
	mapper.AutoMapper(&file, &uploadFileServiceModel)
	uploadFileServiceModel.Extension = strings.Split(file.FileName, ".")[1]
	uploadFileServiceModel.Path = utility.GetRelPath(file.RelationType)
	uploadResult, err := utility.UploadFileService.UploadFile(uploadFileServiceModel)
	if err != nil {
		return nil
	}
	jalaliDate := ptime.Now()
	numericJalaliDate, _ := strconv.Atoi(jalaliDate.Format("yyyyMMdd"))
	fileEntity := Entities.File{
		Id:                uuid.New().String(),
		File_Name:         file.FileName,
		File_Length:       file.FileLength,
		Content_Type:      gomime.TypeByExtension(strings.Split(file.FileName, ".")[1]),
		Extension:         strings.Split(file.FileName, ".")[1],
		Relation_Type:     file.RelationType,
		Relation_Id:       file.RelationId,
		Created_At:        time.Now(),
		Jalali_Created_At: numericJalaliDate,
		Is_Deleted:        false,
		Relative_Url:      uploadResult.RelativeUrl,
		Base_Url:          uploadResult.BaseUrl,
	}
	utility.IFileRepository.Create(&fileEntity)
	return &fileEntity
}
