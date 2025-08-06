package payment

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type GeneralPaymentController struct {
	constants      *bootstrap.Constants
	paymentService usecase.PaymentService
}

func NewGeneralPaymentController(
	constants *bootstrap.Constants,
	paymentService usecase.PaymentService,
) *GeneralPaymentController {
	return &GeneralPaymentController{
		constants:      constants,
		paymentService: paymentService,
	}
}

func (corporationController *GeneralPaymentController) GetPaymentMethods(ctx *gin.Context) {
	methods := corporationController.paymentService.GetPaymentMethods()
	controller.Response(ctx, 200, "", methods)
}
