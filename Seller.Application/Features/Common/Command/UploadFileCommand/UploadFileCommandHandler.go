package UploadFileCommand

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Model"
	"github.com/devfeel/mapper"
)

type UploadFileCommandHandler struct {
	IFileRepository Presistence.IFileRepository
	FileUtility     Common.FileUtility
}

func (handler UploadFileCommandHandler) HandlerFunc(command UploadFileCommand) Common.ResponseModel[UploadFileVm] {

	uploadFileModel := Model.UploadFileUtilityInputDTO{
		//FileBytes:    command.FileBytes,
		//FileName:     command.FileName,
		//FileLength:   command.FileLength,
		//Extension:    filepath.Ext(command.FileName),
		//RelationType: command.RelationType,
	}
	mapper.AutoMapper(&command, &uploadFileModel)
	entity := handler.FileUtility.UploadFile(uploadFileModel)
	result := UploadFileVm{
		FileId: entity.Id,
		Url:    entity.Base_Url + entity.Relative_Url,
	}
	return (Common.ResponseModel[UploadFileVm]{}).ApiResponse(&result, Common.STATUS_SUCCESS)
}
