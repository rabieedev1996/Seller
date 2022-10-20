package AddToBasket

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Infrastructure/Service"
)

type AddToBasketCommandHandler struct {
	RedisService         Service.RedisService
	IProductRepository   Presistence.IProductRepository
	IWarehouseRepository Presistence.IWarehouseRepository
}

func (handler AddToBasketCommandHandler) HandlerFunc(command AddToBasketCommand) Common.ResponseModel[string] {

	return (Common.ResponseModel[string]{}).ApiResponse("", Common.STATUS_SUCCESS)
}
