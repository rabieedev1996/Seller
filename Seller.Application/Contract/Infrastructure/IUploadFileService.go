package Infrastructure

import "Seller/Seller.Application/Model"

type IUploadFileService interface {
	UploadFile(model Model.UploadFileServiceInputDTO) (Model.UploadFileServiceOutputDTO, error)
}
