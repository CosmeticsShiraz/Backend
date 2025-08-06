package httpv1

import (
	"github.com/CosmeticsShiraz/Backend/wire"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(routerGroup *gin.RouterGroup, app *wire.Application) {
	const status string = "/status"

	accessManagement := routerGroup.Group("")
	// accessManagement.Use(app.Middlewares.Auth.RequirePermission([]enums.PermissionType{enums.ManageUsers, enums.ManageRoles}))
	{
		permissions := accessManagement.Group("/permissions")
		{
			permissions.GET("", app.Controllers.Admin.UserController.GetPermissionsList)
			permissions.GET("/:permissionID/roles", app.Controllers.Admin.UserController.GetPermissionRoles)
		}

		roles := accessManagement.Group("/roles")
		{
			roles.GET("", app.Controllers.Admin.UserController.GetRolesList)
			roles.POST("", app.Controllers.Admin.UserController.CreateRole)

			rolesSubGroup := roles.Group("/:roleID")
			{
				rolesSubGroup.GET("", app.Controllers.Admin.UserController.GetRoleDetails)
				rolesSubGroup.GET("/owners", app.Controllers.Admin.UserController.GetRoleOwners)
				rolesSubGroup.PUT("", app.Controllers.Admin.UserController.UpdateRole)
				rolesSubGroup.DELETE("", app.Controllers.Admin.UserController.DeleteRole)
			}
		}

		userRoles := accessManagement.Group("/users/:userID/roles")
		{
			userRoles.GET("", app.Controllers.Admin.UserController.GetUserRoles)
			userRoles.PUT("", app.Controllers.Admin.UserController.UpdateUserRoles)
		}
	}

	userManagement := routerGroup.Group("/users")
	{
		userManagement.GET("", app.Controllers.Admin.UserController.GetUsers)
		userManagement.PUT("/:userID/ban", app.Controllers.Admin.UserController.BanUser)
		userManagement.PUT("/:userID/unban", app.Controllers.Admin.UserController.UnbanUser)
	}

	report := routerGroup.Group("/report")
	{
		report.GET("/maintenance", app.Controllers.Admin.ReportController.GetMaintenanceReports)
		report.GET("/panel", app.Controllers.Admin.ReportController.GetPanelReports)
		report.POST("/resolve/:reportID", app.Controllers.Admin.ReportController.ResolveReport)
	}

	news := routerGroup.Group("/news")
	{
		news.POST("/draft", app.Controllers.Admin.NewsController.CreateDraftNews)
		news.GET("", app.Controllers.Admin.NewsController.GetNewsList)
		news.GET(status, app.Controllers.Admin.NewsController.GetAllNewsStatuses)
		news.DELETE("", app.Controllers.Admin.NewsController.DeleteNews)
		newsSubgroup := news.Group("/:newsID")
		{
			newsSubgroup.GET("", app.Controllers.Admin.NewsController.GetNews)
			newsSubgroup.PUT("", app.Controllers.Admin.NewsController.EditNews)
			newsSubgroup.PUT("/publish", app.Controllers.Admin.NewsController.PublishNews)
			newsSubgroup.PUT("unpublish", app.Controllers.Admin.NewsController.UnpublishNews)
			newsSubgroup.POST("/media", app.Controllers.Admin.NewsController.AddNewsMedia)
			newsSubgroup.DELETE("/media/:mediaID", app.Controllers.Admin.NewsController.DeleteNewsMedia)
			newsSubgroup.GET("/media/:mediaID", app.Controllers.Admin.NewsController.GetNewsMedia)
		}

	}
}
