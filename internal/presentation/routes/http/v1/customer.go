package httpv1

import (
	"github.com/CosmeticsShiraz/Backend/wire"
	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(routerGroup *gin.RouterGroup, app *wire.Application) {
	profile := routerGroup.Group("/profile")
	{
		profile.GET("", app.Controllers.Customer.UserController.GetMyProfile)
		profile.PUT("/password", app.Controllers.Customer.UserController.ResetPassword)
		profile.POST("/complete", app.Controllers.Customer.UserController.CompleteRegister)
		profile.POST("/verify/email", app.Controllers.Customer.UserController.VerifyEmail)
		profile.PUT("", app.Controllers.Customer.UserController.UpdateProfile)
	}

	maintenances := routerGroup.Group("/maintenance/request")
	{
		maintenances.GET("/level", app.Controllers.Customer.MaintenanceController.GetMaintenanceUrgencyLevels)
		maintenances.POST("", app.Controllers.Customer.MaintenanceController.CreateMaintenanceRequest)
		maintenances.GET("", app.Controllers.Customer.MaintenanceController.GetAllMaintenanceRequests)
		requestsSubGroup := maintenances.Group("/:requestID")
		{
			requestsSubGroup.GET("", app.Controllers.Customer.MaintenanceController.GetMaintenanceRequest)
			requestsSubGroup.PUT("", app.Controllers.Customer.MaintenanceController.UpdateMaintenanceRequest)
			requestsSubGroup.PUT("/cancel", app.Controllers.Customer.MaintenanceController.CancelMaintenanceRequest)
			requestsSubGroup.PUT("record/approve", app.Controllers.Customer.MaintenanceController.ApproveMaintenanceRecord)
		}
	}

	addresses := routerGroup.Group("/address")
	{
		addresses.POST("", app.Controllers.Customer.AddressController.CreateUserAddress)
		addresses.GET("", app.Controllers.Customer.AddressController.GetCustomerAddresses)
	}

	chat := routerGroup.Group("/chat")
	{
		chat.GET("/room", app.Controllers.Customer.ChatController.GetUserRooms)
		chat.GET("/room/:roomID/messages", app.Controllers.Customer.ChatController.GetMessages)
		chat.PUT("/room/:roomID/block", app.Controllers.Customer.ChatController.BlockRoom)
		chat.PUT("/room/:roomID/unblock", app.Controllers.Customer.ChatController.UnBlockRoom)
	}

	notification := routerGroup.Group("/notifications")
	{
		notification.POST("/:notificationID/read", app.Controllers.Customer.NotificationController.MarkAsRead)
		notification.GET("", app.Controllers.Customer.NotificationController.GetUserNotifications)
		notification.GET("/setting", app.Controllers.Customer.NotificationController.GetUserNotificationSettings)
		notification.PUT("/setting/:settingID", app.Controllers.Customer.NotificationController.UpdateSettings)
	}

	tickets := routerGroup.Group("/ticket")
	{
		tickets.POST("", app.Controllers.Customer.TicketController.CreateTicket)
		tickets.GET("/list", app.Controllers.Customer.TicketController.GetTickets)
		ticketSubGroup := tickets.Group("/:ticketID/comments")
		{
			ticketSubGroup.GET("", app.Controllers.Customer.TicketController.GetComments)
			ticketSubGroup.POST("", app.Controllers.Customer.TicketController.CreateComment)
		}
	}

	reports := routerGroup.Group("/report")
	{
		maintenanceReports := reports.Group("/maintenance")
		{
			maintenanceReports.POST("/:recordID", app.Controllers.Customer.ReportController.CreateMaintenanceReport)
		}
		panelReports := reports.Group("/panel")
		{
			panelReports.POST("/:panelID", app.Controllers.Customer.ReportController.CreatePanelReport)
		}
	}
}
