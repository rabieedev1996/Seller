package Controller

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Common/Command/UploadFileCommand"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonController struct {
	IFileRepository Presistence.IFileRepository
	FileUtility     Common.FileUtility
	Context         *gin.Context
}

func (controller CommonController) UploadFile() *CommonController {
	header, _ := controller.Context.FormFile("uploadfile")
	var fileBytes []byte
	file, _ := header.Open()
	file.Read(fileBytes)

	handler := UploadFileCommand.UploadFileCommandHandler{
		IFileRepository: controller.IFileRepository,
		FileUtility:     controller.FileUtility,
	}
	uploadFileCommand := UploadFileCommand.UploadFileCommand{
		FileBytes:  fileBytes,
		FileLength: header.Size,
		FileName:   header.Filename,
	}
	bussinesResult := handler.HandlerFunc(uploadFileCommand)
	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}
