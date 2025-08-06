package address

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	addressdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/address"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type CustomerAddressController struct {
	constants      *bootstrap.Constants
	addressService usecase.AddressService
}

func NewCustomerAddressController(
	constants *bootstrap.Constants,
	addressService usecase.AddressService,
) *CustomerAddressController {
	return &CustomerAddressController{
		constants:      constants,
		addressService: addressService,
	}
}

// read users from table name maybe ? or use enums or constants instead ?
func (addressController *CustomerAddressController) CreateUserAddress(ctx *gin.Context) {
	type createAddressParams struct {
		ProvinceID    uint   `json:"provinceID" validate:"required"`
		CityID        uint   `json:"cityID" validate:"required"`
		StreetAddress string `json:"streetAddress" validate:"required"`
		PostalCode    string `json:"postalCode" validate:"required"`
		HouseNumber   string `json:"houseNumber" validate:"required"`
		Unit          uint   `json:"unit" validate:"required"`
	}
	params := controller.Validated[createAddressParams](ctx)
	ownerID, _ := ctx.Get(addressController.constants.Context.ID)
	addressRequestInfo := addressdto.CreateAddressRequest{
		ProvinceID:    params.ProvinceID,
		CityID:        params.CityID,
		StreetAddress: params.StreetAddress,
		PostalCode:    params.PostalCode,
		HouseNumber:   params.HouseNumber,
		Unit:          params.Unit,
		OwnerID:       ownerID.(uint),
		OwnerType:     addressController.constants.AddressOwners.User,
	}
	createdAddress, err := addressController.addressService.CreateAddress(addressRequestInfo)
	if err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, addressController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.createAddress")
	controller.Response(ctx, 200, message, createdAddress)
}

func (addressController *CustomerAddressController) GetCustomerAddresses(ctx *gin.Context) {
	ownerID, _ := ctx.Get(addressController.constants.Context.ID)
	ownerInfo := addressdto.GetOwnerAddressesRequest{
		OwnerID:   ownerID.(uint),
		OwnerType: addressController.constants.AddressOwners.User,
	}
	addresses, err := addressController.addressService.GetAddresses(ownerInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", addresses)
}
