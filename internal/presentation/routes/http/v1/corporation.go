package httpv1

import (
	"github.com/CosmeticsShiraz/Backend/wire"
	"github.com/gin-gonic/gin"
)

func SetupCorporationRoutes(routerGroup *gin.RouterGroup, app *wire.Application) {
	const status string = "/status"

	profile := routerGroup.Group("/:corporationID/profile")
	{
		profile.GET("", app.Controllers.Corporation.CorporationController.GetMyProfile)
		profile.POST("/address", app.Controllers.Corporation.CorporationController.AddAddress)
		profile.DELETE("/address/:addressID", app.Controllers.Corporation.CorporationController.DeleteAddress)
		profile.POST("/contacts", app.Controllers.Corporation.CorporationController.AddContactInformation)
		profile.DELETE("/contacts/:contactID", app.Controllers.Corporation.CorporationController.DeleteContactInformation)
		profile.PUT("/logo", app.Controllers.Corporation.CorporationController.ChangeLogo)
	}

	guarantees := routerGroup.Group("/:corporationID/guarantee")
	{
		guarantees.GET("", app.Controllers.Corporation.GuaranteeController.GetGuarantees)
		guarantees.GET("/type", app.Controllers.Corporation.GuaranteeController.GetGuaranteeTypes)
		guarantees.POST("", app.Controllers.Corporation.GuaranteeController.CreateGuarantee)
		guaranteesSubGroup := guarantees.Group("/:guaranteeID")
		{
			guaranteesSubGroup.GET("", app.Controllers.Corporation.GuaranteeController.GetGuarantee)
			guaranteesSubGroup.PUT(status, app.Controllers.Corporation.GuaranteeController.UpdateGuarantee)
		}
	}

	maintenances := routerGroup.Group("/:corporationID/maintenance")
	{
		requests := maintenances.Group("/request")
		{
			requests.GET(status, app.Controllers.Corporation.MaintenanceController.GetMaintenanceStatuses)
			requests.GET("", app.Controllers.Corporation.MaintenanceController.GetAllMaintenanceRequests)
			requestsSubGroup := requests.Group("/:requestID")
			{
				requestsSubGroup.GET("", app.Controllers.Corporation.MaintenanceController.GetMaintenanceRequest)
				requestsSubGroup.PUT("/accept", app.Controllers.Corporation.MaintenanceController.AcceptMaintenanceRequest)
				requestsSubGroup.PUT("/reject", app.Controllers.Corporation.MaintenanceController.RejectMaintenanceRequest)
				records := requestsSubGroup.Group("/record")
				{
					records.POST("", app.Controllers.Corporation.MaintenanceController.CreateMaintenanceRecord)
					records.PUT("", app.Controllers.Corporation.MaintenanceController.UpdateMaintenanceRecord)
				}
			}
		}
	}

	chat := routerGroup.Group("/chat")
	{
		chat.GET("/room/:corporationID", app.Controllers.Corporation.ChatController.GetRoom)
		chat.GET("/rooms/:corporationID", app.Controllers.Corporation.ChatController.GetRooms)
		chat.PUT("/room/:roomID/block", app.Controllers.Corporation.ChatController.BlockRoom)
		chat.PUT("/room/:roomID/unblock", app.Controllers.Corporation.ChatController.UnBlockRoom)
	}
}
