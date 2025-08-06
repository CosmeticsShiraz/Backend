package maintenance

import (
	bootstrap "github.com/CosmeticsShiraz/Backend/bootstrap"

	guaranteedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/guarantee"
	maintenancedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/maintenance"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type CorporationMaintenanceController struct {
	constants          *bootstrap.Constants
	pagination         *bootstrap.Pagination
	maintenanceService usecase.MaintenanceService
}

func NewCorporationMaintenanceController(
	constants *bootstrap.Constants,
	pagination *bootstrap.Pagination,
	maintenanceService usecase.MaintenanceService,
) *CorporationMaintenanceController {
	return &CorporationMaintenanceController{
		constants:          constants,
		pagination:         pagination,
		maintenanceService: maintenanceService,
	}
}

func (maintenanceController *CorporationMaintenanceController) GetMaintenanceStatuses(ctx *gin.Context) {
	statuses := maintenanceController.maintenanceService.GetMaintenanceRequestStatuses(enum.AgentTypeCorporation)

	controller.Response(ctx, 200, "", statuses)
}

func (maintenanceController *CorporationMaintenanceController) GetAllMaintenanceRequests(ctx *gin.Context) {
	type maintenanceRequestsParams struct {
		CorporationID uint `uri:"corporationID" validate:"required"`
		Status        uint `form:"status" validate:"required"`
	}
	params := controller.Validated[maintenanceRequestsParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	pagination := controller.GetPagination(ctx, maintenanceController.pagination.DefaultPage, maintenanceController.pagination.DefaultPageSize)
	offset, limit := pagination.GetOffsetLimit()

	listInfo := maintenancedto.CorporationMaintenanceListRequest{
		CorporationID: params.CorporationID,
		OperatorID:    operatorID.(uint),
		Status:        params.Status,
		Offset:        offset,
		Limit:         limit,
	}
	requests, err := maintenanceController.maintenanceService.GetCorporationMaintenanceRequests(listInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", requests)
}

func (maintenanceController *CorporationMaintenanceController) GetMaintenanceRequest(ctx *gin.Context) {
	type maintenanceRequestParams struct {
		CorporationID uint `uri:"corporationID" validate:"required"`
		RequestID     uint `uri:"requestID" validate:"required"`
	}
	params := controller.Validated[maintenanceRequestParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	listInfo := maintenancedto.CorporationMaintenanceRequest{
		CorporationID: params.CorporationID,
		OperatorID:    operatorID.(uint),
		RequestID:     params.RequestID,
	}
	requests, err := maintenanceController.maintenanceService.GetCorporationMaintenanceRequest(listInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", requests)
}

func (maintenanceController *CorporationMaintenanceController) AcceptMaintenanceRequest(ctx *gin.Context) {
	type maintenanceRequestParams struct {
		CorporationID uint `uri:"corporationID" validate:"required"`
		RequestID     uint `uri:"requestID" validate:"required"`
	}
	params := controller.Validated[maintenanceRequestParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	requestInfo := maintenancedto.CorporationMaintenanceRequest{
		CorporationID: params.CorporationID,
		OperatorID:    operatorID.(uint),
		RequestID:     params.RequestID,
	}
	if err := maintenanceController.maintenanceService.AcceptMaintenanceRequest(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, maintenanceController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.acceptMaintenanceRequest")
	controller.Response(ctx, 200, message, nil)
}

func (maintenanceController *CorporationMaintenanceController) RejectMaintenanceRequest(ctx *gin.Context) {
	type maintenanceRequestParams struct {
		CorporationID uint `uri:"corporationID" validate:"required"`
		RequestID     uint `uri:"requestID" validate:"required"`
	}
	params := controller.Validated[maintenanceRequestParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	requestInfo := maintenancedto.CorporationMaintenanceRequest{
		CorporationID: params.CorporationID,
		OperatorID:    operatorID.(uint),
		RequestID:     params.RequestID,
	}
	if err := maintenanceController.maintenanceService.RejectMaintenanceRequest(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, maintenanceController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.rejectMaintenanceRequest")
	controller.Response(ctx, 200, message, nil)
}

func (maintenanceController *CorporationMaintenanceController) CreateMaintenanceRecord(ctx *gin.Context) {
	type guaranteeViolation struct {
		Reason  string `json:"reason" validate:"required"`
		Details string `json:"details" validate:"required"`
	}
	type createRecordParams struct {
		RequestID          uint                `uri:"requestID" validate:"required"`
		CorporationID      uint                `uri:"corporationID" validate:"required"`
		Title              string              `json:"title" validate:"required"`
		Details            string              `json:"details" validate:"required"`
		GuaranteeViolation *guaranteeViolation `json:"guaranteeViolation"`
	}
	params := controller.Validated[createRecordParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	var guaranteeViolationParams *guaranteedto.CreateGuaranteeViolationRequest = nil
	if params.GuaranteeViolation != nil {
		guaranteeViolationParams = &guaranteedto.CreateGuaranteeViolationRequest{
			CorporationID: params.CorporationID,
			OperatorID:    operatorID.(uint),
			Reason:        params.GuaranteeViolation.Reason,
			Details:       params.GuaranteeViolation.Details,
		}
	}

	recordInfo := maintenancedto.CreateMaintenanceRecordRequest{
		CorporationID:      params.CorporationID,
		OperatorID:         operatorID.(uint),
		RequestID:          params.RequestID,
		Title:              params.Title,
		Details:            params.Details,
		GuaranteeViolation: guaranteeViolationParams,
	}
	if err := maintenanceController.maintenanceService.CreateMaintenanceRecord(recordInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, maintenanceController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.addMaintenanceRecord")
	controller.Response(ctx, 200, message, nil)
}

func (maintenanceController *CorporationMaintenanceController) UpdateMaintenanceRecord(ctx *gin.Context) {
	type guaranteeViolation struct {
		Reason  *string `json:"reason"`
		Details *string `json:"details"`
	}
	type createRecordParams struct {
		RequestID          uint                `uri:"requestID" validate:"required"`
		CorporationID      uint                `uri:"corporationID" validate:"required"`
		Title              *string             `json:"title"`
		Details            *string             `json:"details"`
		GuaranteeViolation *guaranteeViolation `json:"guaranteeViolation"`
	}
	params := controller.Validated[createRecordParams](ctx)
	operatorID, _ := ctx.Get(maintenanceController.constants.Context.ID)

	var guaranteeViolationParams *guaranteedto.UpdateGuaranteeViolationRequest = nil
	if params.GuaranteeViolation != nil {
		guaranteeViolationParams = &guaranteedto.UpdateGuaranteeViolationRequest{
			OperatorID: operatorID.(uint),
			Reason:     params.GuaranteeViolation.Reason,
			Details:    params.GuaranteeViolation.Details,
		}
	}

	recordInfo := maintenancedto.UpdateMaintenanceRecordRequest{
		CorporationID:      params.CorporationID,
		OperatorID:         operatorID.(uint),
		RequestID:          params.RequestID,
		Title:              params.Title,
		Details:            params.Details,
		GuaranteeViolation: guaranteeViolationParams,
	}
	if err := maintenanceController.maintenanceService.UpdateMaintenanceRecord(recordInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, maintenanceController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.updateMaintenanceRecord")
	controller.Response(ctx, 200, message, nil)
}
