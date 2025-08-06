package address

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	addressdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/address"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type GeneralAddressController struct {
	constants      *bootstrap.Constants
	addressService usecase.AddressService
}

func NewGeneralAddressController(
	constants *bootstrap.Constants,
	addressService usecase.AddressService,
) *GeneralAddressController {
	return &GeneralAddressController{
		constants:      constants,
		addressService: addressService,
	}
}

func (addressController *GeneralAddressController) GetProvince(ctx *gin.Context) {
	provinceList, err := addressController.addressService.GetProvinceList()
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", provinceList)
}

func (addressController *GeneralAddressController) GetProvinceCities(ctx *gin.Context) {
	type getCitiesParams struct {
		ProvinceID uint `uri:"provinceID" validate:"required"`
	}
	params := controller.Validated[getCitiesParams](ctx)
	provinceInfo := addressdto.GetProvinceCitiesRequest{
		ProvinceID: params.ProvinceID,
	}
	citiesList, err := addressController.addressService.GetCityProvinceCities(provinceInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", citiesList)
}
