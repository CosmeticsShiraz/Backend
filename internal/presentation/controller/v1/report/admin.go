package report

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	reportdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/report"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type AdminReportController struct {
	constants     *bootstrap.Constants
	pagination    *bootstrap.Pagination
	reportService usecase.ReportService
}

func NewAdminReportController(
	constants *bootstrap.Constants,
	pagination *bootstrap.Pagination,
	reportService usecase.ReportService,
) *AdminReportController {
	return &AdminReportController{
		constants:     constants,
		pagination:    pagination,
		reportService: reportService,
	}
}

func (reportController *AdminReportController) GetMaintenanceReports(ctx *gin.Context) {
	type GetMaintenanceReportsRequest struct {
		Status uint `form:"status" validate:"required"`
	}
	params := controller.Validated[GetMaintenanceReportsRequest](ctx)
	pagination := controller.GetPagination(ctx, reportController.pagination.DefaultPage, reportController.pagination.DefaultPageSize)
	offset, limit := pagination.GetOffsetLimit()
	ownerID, _ := ctx.Get(reportController.constants.Context.ID)

	requestInfo := reportdto.ReportListRequest{
		OwnerID: ownerID.(uint),
		Status:  params.Status,
		Offset:  offset,
		Limit:   limit,
	}
	reports, err := reportController.reportService.GetMaintenanceReports(requestInfo)
	if err != nil {
		panic(err)
	}
	controller.Response(ctx, 200, "", reports)
}

func (reportController *AdminReportController) GetPanelReports(ctx *gin.Context) {
	type GetPanelReportsRequest struct {
		Status uint `form:"status" validate:"required"`
	}
	params := controller.Validated[GetPanelReportsRequest](ctx)
	pagination := controller.GetPagination(ctx, reportController.pagination.DefaultPage, reportController.pagination.DefaultPageSize)
	offset, limit := pagination.GetOffsetLimit()
	ownerID, _ := ctx.Get(reportController.constants.Context.ID)

	requestInfo := reportdto.ReportListRequest{
		OwnerID: ownerID.(uint),
		Status:  params.Status,
		Offset:  offset,
		Limit:   limit,
	}
	reports, err := reportController.reportService.GetPanelReports(requestInfo)
	if err != nil {
		panic(err)
	}
	controller.Response(ctx, 200, "", reports)
}

func (reportController *AdminReportController) ResolveReport(ctx *gin.Context) {
	type ResolveReportRequest struct {
		ReportID uint `uri:"reportID" validate:"required"`
	}
	params := controller.Validated[ResolveReportRequest](ctx)
	userID, _ := ctx.Get(reportController.constants.Context.ID)

	requestInfo := reportdto.ResolveReportRequest{
		ReportID: params.ReportID,
		UserID:   userID.(uint),
	}
	if err := reportController.reportService.ResolveReport(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, reportController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.reportResolved")
	controller.Response(ctx, 200, message, nil)
}
