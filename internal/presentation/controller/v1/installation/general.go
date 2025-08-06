package installation

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type GeneralInstallationController struct {
	constants           *bootstrap.Constants
	installationService usecase.InstallationService
}

func NewGeneralInstallationController(
	constants *bootstrap.Constants,
	installationService usecase.InstallationService,
) *GeneralInstallationController {
	return &GeneralInstallationController{
		constants:           constants,
		installationService: installationService,
	}
}

func (installationController *GeneralInstallationController) GetRequestStatuses(ctx *gin.Context) {
	statuses := installationController.installationService.GetRequestStatuses()
	controller.Response(ctx, 200, "", statuses)
}

func (installationController *GeneralInstallationController) GetPanelStatuses(ctx *gin.Context) {
	statuses := installationController.installationService.GetPanelStatuses()
	controller.Response(ctx, 200, "", statuses)
}

func (installationController *GeneralInstallationController) GetBuildingTypes(ctx *gin.Context) {
	types := installationController.installationService.GetBuildingTypes()
	controller.Response(ctx, 200, "", types)
}
