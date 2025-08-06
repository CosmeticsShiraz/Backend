package report

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	reportdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/report"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type CustomerReportController struct {
	constants     *bootstrap.Constants
	reportService usecase.ReportService
}

func NewCustomerReportController(
	constants *bootstrap.Constants,
	reportService usecase.ReportService,
) *CustomerReportController {
	return &CustomerReportController{
		constants:     constants,
		reportService: reportService,
	}
}

func (reportController *CustomerReportController) CreateMaintenanceReport(ctx *gin.Context) {
	type createMaintenanceReportRequest struct {
		RecordID    uint   `uri:"recordID" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	params := controller.Validated[createMaintenanceReportRequest](ctx)
	userID, _ := ctx.Get(reportController.constants.Context.ID)

	requestInfo := reportdto.CreateReportRequest{
		ObjectID:       params.RecordID,
		ObjectType:     reportController.constants.ReportObjectTypes.Maintenance,
		ReportedByID:   userID.(uint),
		ReportedByType: reportController.constants.ReportOwners.User,
		Description:    params.Description,
	}
	if err := reportController.reportService.CreateMaintenanceReport(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, reportController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.createReport")
	controller.Response(ctx, 200, message, nil)
}

func (reportController *CustomerReportController) CreatePanelReport(ctx *gin.Context) {
	type createPanelReportRequest struct {
		PanelID     uint   `uri:"panelID" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	params := controller.Validated[createPanelReportRequest](ctx)
	userID, _ := ctx.Get(reportController.constants.Context.ID)

	requestInfo := reportdto.CreateReportRequest{
		ObjectID:       params.PanelID,
		ObjectType:     reportController.constants.ReportObjectTypes.Panel,
		ReportedByID:   userID.(uint),
		ReportedByType: reportController.constants.ReportOwners.User,
		Description:    params.Description,
	}
	if err := reportController.reportService.CreatePanelReport(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, reportController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.createReport")
	controller.Response(ctx, 200, message, nil)
}
