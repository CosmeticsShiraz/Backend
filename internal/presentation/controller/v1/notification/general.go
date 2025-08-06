package notification

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type GeneralNotificationController struct {
	constants           *bootstrap.Constants
	notificationService usecase.NotificationService
}

func NewGeneralNotificationController(
	constants *bootstrap.Constants,
	notificationService usecase.NotificationService,
) *GeneralNotificationController {
	return &GeneralNotificationController{
		constants:           constants,
		notificationService: notificationService,
	}
}

func (notificationController *GeneralNotificationController) GetContactTypes(ctx *gin.Context) {
	notificationTypes, err := notificationController.notificationService.GetNotificationsType()
	if err != nil {
		panic(err)
	}
	controller.Response(ctx, 200, "", notificationTypes)
}
