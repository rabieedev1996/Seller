package GinHandler

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"Seller/Seller.Infrastructure/Service"
	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	Engine *gin.RouterGroup
}

// @Summary Upload Single File
// @Tags Common
// @Success 200 {object} Common.ResponseModel[string]
// @Router /Common/UploadFile [post]
// @Security bearer
func (hanlder CommonHandler) UploadFile() *CommonHandler {
	hanlder.Engine.POST("/UploadFile", func(c *gin.Context) {
		hanlder.createController(c).UploadFile()
	})
	return &hanlder
}

func (hanlder CommonHandler) createController(c *gin.Context) Controller.CommonController {
	db := Config.GetDatabaseConnection()

	var uploadFileService Infrastructure.IUploadFileService
	uploadFileService = Service.ArvanCloudUploadFileService{}
	fileUtility := Common.FileUtility{
		UploadFileService: uploadFileService,
		IFileRepository:   nil,
	}
	controller := Controller.CommonController{
		Context: c,
		IFileRepository: DBRepositories.FileRepository{
			Generic: Common.GetGenericRepository[Entities.File](db),
		},
		FileUtility: fileUtility,
	}
	return controller
}
