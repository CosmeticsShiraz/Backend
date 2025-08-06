package httpv1

import (
	"github.com/CosmeticsShiraz/Backend/wire"
	"github.com/gin-gonic/gin"
)

func SetupGeneralRoutes(routerGroup *gin.RouterGroup, app *wire.Application) {
	const status string = "/status"

	auth := routerGroup.Group("/auth")
	{
		auth.POST("/register/basic", app.Controllers.General.UserController.BasicRegister)
		auth.POST("/verify/phone", app.Controllers.General.UserController.VerifyPhone)
		auth.POST("/login", app.Controllers.General.UserController.Login)
		auth.POST("/forgot-password", app.Controllers.General.UserController.ForgotPassword)
		auth.POST("/confirm-otp", app.Controllers.General.UserController.ConfirmOTP)
		auth.POST("/refresh", app.Controllers.General.UserController.RefreshToken)
	}

	addresses := routerGroup.Group("/address")
	{
		addresses.GET("/province", app.Controllers.General.AddressController.GetProvince)
		addresses.GET("/province/:provinceID/city", app.Controllers.General.AddressController.GetProvinceCities)
	}

	contacts := routerGroup.Group("/contact")
	{
		contacts.GET("/types", app.Controllers.General.CorporationController.GetContactTypes)
	}

	corporations := routerGroup.Group("/corporation")
	{
		corporations.GET("", app.Controllers.General.CorporationController.GetCorporations)
	}

	notifications := routerGroup.Group("/notifications")
	{
		notifications.GET("/type", app.Controllers.General.NotificationController.GetContactTypes)
	}

	guarantees := routerGroup.Group("/guarantee")
	{
		guarantees.GET(status, app.Controllers.Corporation.GuaranteeController.GetGuaranteeStatuses)
	}

	maintenances := routerGroup.Group("/maintenance")
	{
		maintenances.GET(status, app.Controllers.Customer.MaintenanceController.GetMaintenanceStatuses)
	}

	payments := routerGroup.Group("/payment")
	{
		payments.GET("method", app.Controllers.General.PaymentController.GetPaymentMethods)
	}

	news := routerGroup.Group("/news")
	{
		news.GET("", app.Controllers.General.NewsController.GetNewsList)
		news.GET("/:newsID", app.Controllers.General.NewsController.GetNews)
		news.GET("/:newsID/media/:mediaID", app.Controllers.General.NewsController.GetNewsMedia)
	}
}
