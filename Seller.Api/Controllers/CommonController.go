package Controller

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Common/Command/UploadFileCommand"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type CommonController struct {
	IFileRepository Presistence.IFileRepository
	FileUtility     Common.FileUtility
	Context         *gin.Context
}

func (controller CommonController) UploadFile() *CommonController {

	file, header, err := controller.Context.Request.FormFile("file")
	//defer file.Close()
	if err != nil {
		result := (Common.ResponseModel[UploadFileCommand.UploadFileVm]{}).ApiResponse(nil, Common.STATUS_FILE_EMPETY)
		controller.Context.JSON(http.StatusOK, result)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		result := (Common.ResponseModel[UploadFileCommand.UploadFileVm]{}).ApiResponse(nil, Common.STATUS_FILE_EMPETY)
		controller.Context.JSON(http.StatusOK, result)
	}
	uploadFileCommand := UploadFileCommand.UploadFileCommand{
		FileBytes:  buf.Bytes(),
		FileLength: header.Size,
		FileName:   header.Filename,
	}
	handler := UploadFileCommand.UploadFileCommandHandler{
		IFileRepository: controller.IFileRepository,
		FileUtility:     controller.FileUtility,
	}

	bussinesResult := handler.HandlerFunc(uploadFileCommand)
	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}
