package service

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	installationdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/installation"
	maintenancedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/maintenance"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/domain/exception"
	"github.com/CosmeticsShiraz/Backend/internal/domain/repository/postgres"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
	postgresImpl "github.com/CosmeticsShiraz/Backend/internal/infrastructure/repository/postgres"
)

type MaintenanceService struct {
	constants             *bootstrap.Constants
	userService           usecase.UserService
	installationService   usecase.InstallationService
	corporationService    usecase.CorporationService
	guaranteeService      usecase.GuaranteeService
	maintenanceRepository postgres.MaintenanceRepository
	db                    database.Database
}

func NewMaintenanceService(
	constants *bootstrap.Constants,
	userService usecase.UserService,
	installationService usecase.InstallationService,
	corporationService usecase.CorporationService,
	guaranteeService usecase.GuaranteeService,
	maintenanceRepository postgres.MaintenanceRepository,
	db database.Database,
) *MaintenanceService {
	return &MaintenanceService{
		constants:             constants,
		userService:           userService,
		installationService:   installationService,
		corporationService:    corporationService,
		guaranteeService:      guaranteeService,
		maintenanceRepository: maintenanceRepository,
		db:                    db,
	}
}

func (maintenanceService *MaintenanceService) getRequest(requestID uint) (*entity.MaintenanceRequest, error) {
	maintenanceRequest, err := maintenanceService.maintenanceRepository.FindRequestByID(maintenanceService.db, requestID)
	if err != nil {
		return nil, err
	}
	if maintenanceRequest == nil {
		notFoundError := exception.NotFoundError{Item: maintenanceService.constants.Field.MaintenanceRequest}
		return nil, notFoundError
	}
	return maintenanceRequest, nil
}

func (maintenanceService *MaintenanceService) getCorporationRequest(requestID, corporationID uint) (*entity.MaintenanceRequest, error) {
	allowedStatuses := enum.GetAllowedMaintenanceRequestStatuses(enum.AgentTypeCorporation)
	maintenanceRequest, err := maintenanceService.maintenanceRepository.FindCorporationRequestByStatus(maintenanceService.db, requestID, corporationID, allowedStatuses)
	if err != nil {
		return nil, err
	}
	if maintenanceRequest == nil {
		notFoundError := exception.NotFoundError{Item: maintenanceService.constants.Field.MaintenanceRequest}
		return nil, notFoundError
	}
	return maintenanceRequest, nil
}

func (maintenanceService *MaintenanceService) getRecord(recordID uint) (*entity.MaintenanceRecord, error) {
	maintenanceRecord, err := maintenanceService.maintenanceRepository.FindRecordByID(maintenanceService.db, recordID)
	if err != nil {
		return nil, err
	}
	if maintenanceRecord == nil {
		notFoundError := exception.NotFoundError{Item: maintenanceService.constants.Field.MaintenanceRecord}
		return nil, notFoundError
	}
	return maintenanceRecord, nil
}

func (maintenanceService *MaintenanceService) getRequestRecord(requestID uint) (*entity.MaintenanceRecord, error) {
	maintenanceRecord, err := maintenanceService.maintenanceRepository.FindRecordByRequestID(maintenanceService.db, requestID)
	if err != nil {
		return nil, err
	}
	if maintenanceRecord == nil {
		notFoundError := exception.NotFoundError{Item: maintenanceService.constants.Field.MaintenanceRecord}
		return nil, notFoundError
	}
	return maintenanceRecord, nil
}

func (maintenanceService *MaintenanceService) ValidateCustomerRecord(recordID, userID uint) error {
	maintenanceRecord, err := maintenanceService.getRecord(recordID)
	if err != nil {
		return err
	}

	maintenanceRequest, err := maintenanceService.getRequest(maintenanceRecord.RequestID)
	if err != nil {
		return err
	}

	if _, err = maintenanceService.installationService.ValidatePanelOwnership(maintenanceRequest.PanelID, userID); err != nil {
		return err
	}

	return nil
}

func (maintenanceService *MaintenanceService) GetRequestByAdmin(recordID uint) (maintenancedto.AdminMaintenanceRequestResponse, error) {
	var response maintenancedto.AdminMaintenanceRequestResponse
	maintenanceRecord, err := maintenanceService.getRecord(recordID)
	if err != nil {
		return response, err
	}

	maintenanceRequest, err := maintenanceService.getRequest(maintenanceRecord.RequestID)
	if err != nil {
		return response, err
	}

	panel, err := maintenanceService.installationService.GetPanelByAdmin(maintenanceRequest.PanelID)
	if err != nil {
		return response, err
	}

	corporation, err := maintenanceService.corporationService.GetCorporationCredentials(maintenanceRequest.CorporationID)
	if err != nil {
		return response, err
	}

	record, err := maintenanceService.getCorporationMaintenanceRecord(maintenanceRequest.ID, maintenanceRequest.PanelID)
	if err != nil {
		return response, err
	}

	response = maintenancedto.AdminMaintenanceRequestResponse{
		ID:                   maintenanceRequest.ID,
		CreatedAt:            maintenanceRequest.CreatedAt,
		Panel:                panel,
		Corporation:          corporation,
		Subject:              maintenanceRequest.Subject,
		Description:          maintenanceRequest.Description,
		UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
		Status:               maintenanceRequest.Status.String(),
		IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		Record:               record,
	}

	return response, nil
}

func (maintenanceService *MaintenanceService) mapStatusForRole(statusID uint, agent enum.AgentType) []enum.MaintenanceRequestStatus {
	status := enum.MaintenanceRequestStatus(statusID)

	allowedStatuses := enum.GetAllowedMaintenanceRequestStatuses(agent)

	for _, allowedStatus := range allowedStatuses {
		if status == allowedStatus {
			if status == enum.MaintenanceRequestStatusAll {
				return allowedStatuses
			}
			return []enum.MaintenanceRequestStatus{status}
		}
	}
	return allowedStatuses
}

func (maintenanceService *MaintenanceService) GetMaintenanceUrgencyLevels() []maintenancedto.MaintenanceStatusesResponse {
	levels := enum.GetAllUrgencyLevels()
	response := make([]maintenancedto.MaintenanceStatusesResponse, len(levels))
	for i, level := range levels {
		response[i] = maintenancedto.MaintenanceStatusesResponse{
			ID:   uint(level),
			Name: level.String(),
		}
	}
	return response
}

func (maintenanceService *MaintenanceService) GetMaintenanceRequestStatuses(agent enum.AgentType) []maintenancedto.MaintenanceStatusesResponse {
	statuses := enum.GetAllowedMaintenanceRequestStatuses(agent)
	response := make([]maintenancedto.MaintenanceStatusesResponse, len(statuses))
	for i, status := range statuses {
		response[i] = maintenancedto.MaintenanceStatusesResponse{
			ID:   uint(status),
			Name: status.String(),
		}
	}
	return response
}

func (maintenanceService *MaintenanceService) CreateMaintenanceRequest(request maintenancedto.CreateMaintenanceRequest) error {
	if err := maintenanceService.userService.IsUserActive(request.OwnerID); err != nil {
		return err
	}

	if err := maintenanceService.corporationService.DoesCorporationExist(request.CorporationID); err != nil {
		return err
	}

	if _, err := maintenanceService.installationService.ValidatePanelOwnership(request.PanelID, request.OwnerID); err != nil {
		return err
	}

	allowedStatus := []enum.MaintenanceRequestStatus{enum.MaintenanceRequestStatusPending}
	currentActiveRequest, err := maintenanceService.maintenanceRepository.FindRequestsByPanelIDAndStatus(maintenanceService.db, request.PanelID, allowedStatus)
	if err != nil {
		return err
	}
	if len(currentActiveRequest) > 0 {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.Pending)
		return conflictErrors
	}

	if request.IsUsingGuarantee {
		if err := maintenanceService.installationService.ValidatePanelGuarantee(request.PanelID); err != nil {
			return err
		}
	}

	maintenanceRequest := &entity.MaintenanceRequest{
		CorporationID:        request.CorporationID,
		PanelID:              request.PanelID,
		Subject:              request.Subject,
		Description:          request.Description,
		Status:               enum.MaintenanceRequestStatusPending,
		UrgencyLevel:         request.UrgencyLevel,
		IsGuaranteeRequested: request.IsUsingGuarantee,
	}
	err = maintenanceService.maintenanceRepository.CreateMaintenanceRequest(maintenanceService.db, maintenanceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (maintenanceService *MaintenanceService) GetCustomerMaintenanceRequests(listInfo maintenancedto.CustomerMaintenanceListRequest) ([]maintenancedto.CustomerMaintenanceRequestResponse, error) {
	allowedStatus := maintenanceService.mapStatusForRole(listInfo.Status, enum.AgentTypeCustomer)

	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	maintenanceRequests, err := maintenanceService.maintenanceRepository.FindRequestsByCustomerID(maintenanceService.db, listInfo.OwnerID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]maintenancedto.CustomerMaintenanceRequestResponse, len(maintenanceRequests))

	for i, maintenanceRequest := range maintenanceRequests {
		panelInfoRequest := installationdto.GetOwnerRequest{
			OwnerID:        listInfo.OwnerID,
			InstallationID: maintenanceRequest.PanelID,
		}
		panel, err := maintenanceService.installationService.GetCustomerPanel(panelInfoRequest)
		if err != nil {
			return nil, err
		}

		corporation, err := maintenanceService.corporationService.GetCorporationCredentials(maintenanceRequest.CorporationID)
		if err != nil {
			return nil, err
		}

		response[i] = maintenancedto.CustomerMaintenanceRequestResponse{
			ID:                   maintenanceRequest.ID,
			CreatedAt:            maintenanceRequest.CreatedAt,
			Panel:                panel,
			Corporation:          corporation,
			Subject:              maintenanceRequest.Subject,
			Description:          maintenanceRequest.Description,
			UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
			Status:               maintenanceRequest.Status.String(),
			IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		}
	}
	return response, nil
}

func (maintenanceService *MaintenanceService) GetCustomerPanelMaintenanceRequests(listInfo maintenancedto.CustomerPanelMaintenanceListRequest) ([]maintenancedto.CustomerMaintenanceRequestResponse, error) {
	if _, err := maintenanceService.installationService.ValidatePanelOwnership(listInfo.PanelID, listInfo.OwnerID); err != nil {
		return nil, err
	}

	allowedStatus := maintenanceService.mapStatusForRole(listInfo.Status, enum.AgentTypeCustomer)

	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	maintenanceRequests, err := maintenanceService.maintenanceRepository.FindRequestsByPanelIDAndStatus(maintenanceService.db, listInfo.PanelID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]maintenancedto.CustomerMaintenanceRequestResponse, len(maintenanceRequests))

	for i, maintenanceRequest := range maintenanceRequests {
		panelInfoRequest := installationdto.GetOwnerRequest{
			OwnerID:        listInfo.OwnerID,
			InstallationID: maintenanceRequest.PanelID,
		}
		panel, err := maintenanceService.installationService.GetCustomerPanel(panelInfoRequest)
		if err != nil {
			return nil, err
		}

		corporation, err := maintenanceService.corporationService.GetCorporationCredentials(maintenanceRequest.CorporationID)
		if err != nil {
			return nil, err
		}

		response[i] = maintenancedto.CustomerMaintenanceRequestResponse{
			ID:                   maintenanceRequest.ID,
			CreatedAt:            maintenanceRequest.CreatedAt,
			Panel:                panel,
			Corporation:          corporation,
			Subject:              maintenanceRequest.Subject,
			Description:          maintenanceRequest.Description,
			UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
			Status:               maintenanceRequest.Status.String(),
			IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		}
	}
	return response, nil
}

func (maintenanceService *MaintenanceService) GetCustomerMaintenanceRequest(maintenanceInfo maintenancedto.CustomerMaintenanceRequest) (maintenancedto.CustomerMaintenanceRequestResponse, error) {
	maintenanceRequest, err := maintenanceService.getRequest(maintenanceInfo.RequestID)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRequestResponse{}, err
	}

	panelInfoRequest := installationdto.GetOwnerRequest{
		OwnerID:        maintenanceInfo.OwnerID,
		InstallationID: maintenanceRequest.PanelID,
	}
	panel, err := maintenanceService.installationService.GetCustomerPanel(panelInfoRequest)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRequestResponse{}, err
	}

	corporation, err := maintenanceService.corporationService.GetCorporationCredentials(maintenanceRequest.CorporationID)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRequestResponse{}, err
	}

	record, err := maintenanceService.getCustomerMaintenanceRecord(maintenanceInfo.RequestID, maintenanceRequest.PanelID)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRequestResponse{}, err
	}

	response := maintenancedto.CustomerMaintenanceRequestResponse{
		ID:                   maintenanceRequest.ID,
		CreatedAt:            maintenanceRequest.CreatedAt,
		Panel:                panel,
		Corporation:          corporation,
		Subject:              maintenanceRequest.Subject,
		Description:          maintenanceRequest.Description,
		UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
		Status:               maintenanceRequest.Status.String(),
		IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		Record:               record,
	}
	return response, nil
}

func (maintenanceService *MaintenanceService) getCustomerMaintenanceRecord(requestID, panelID uint) (maintenancedto.CustomerMaintenanceRecordResponse, error) {
	record, err := maintenanceService.getRequestRecord(requestID)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRecordResponse{}, err
	}

	violation, err := maintenanceService.guaranteeService.GetCustomerPanelGuaranteeViolation(panelID)
	if err != nil {
		return maintenancedto.CustomerMaintenanceRecordResponse{}, err
	}

	recordResponse := maintenancedto.CustomerMaintenanceRecordResponse{
		ID:                 record.ID,
		CreatedAt:          record.CreatedAt,
		Title:              record.Title,
		Details:            record.Details,
		Date:               record.CreatedAt,
		IsUserApproved:     record.IsUserApproved,
		GuaranteeViolation: violation,
	}
	return recordResponse, nil
}

func (maintenanceService *MaintenanceService) UpdateMaintenanceRequest(updateRequest maintenancedto.UpdateCustomerRequest) error {
	maintenanceRequest, err := maintenanceService.getRequest(updateRequest.RequestID)
	if err != nil {
		return err
	}

	if _, err := maintenanceService.installationService.ValidatePanelOwnership(maintenanceRequest.PanelID, updateRequest.OwnerID); err != nil {
		return err
	}

	if maintenanceRequest.Status != enum.MaintenanceRequestStatusPending {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.NotActive)
		return conflictErrors
	}

	if updateRequest.Subject != nil {
		maintenanceRequest.Subject = *updateRequest.Subject
	}

	if updateRequest.Description != nil {
		maintenanceRequest.Description = *updateRequest.Description
	}

	if updateRequest.UrgencyLevel != nil {
		maintenanceRequest.UrgencyLevel = enum.UrgencyLevel(*updateRequest.UrgencyLevel)
	}

	if updateRequest.IsUsingGuarantee != nil {
		if *updateRequest.IsUsingGuarantee {
			if err := maintenanceService.installationService.ValidatePanelGuarantee(maintenanceRequest.PanelID); err != nil {
				return err
			}
			maintenanceRequest.IsGuaranteeRequested = true
		} else {
			maintenanceRequest.IsGuaranteeRequested = false
		}

	}

	if err = maintenanceService.maintenanceRepository.UpdateMaintenanceRequest(maintenanceService.db, maintenanceRequest); err != nil {
		return err
	}
	return nil
}

func (maintenanceService *MaintenanceService) CancelMaintenanceRequest(maintenanceInfo maintenancedto.CustomerMaintenanceRequest) error {
	maintenanceRequest, err := maintenanceService.getRequest(maintenanceInfo.RequestID)
	if err != nil {
		return err
	}

	_, err = maintenanceService.installationService.ValidatePanelOwnership(maintenanceRequest.PanelID, maintenanceInfo.OwnerID)
	if err != nil {
		return err
	}

	var conflictErrors exception.ConflictErrors
	if maintenanceRequest.Status == enum.MaintenanceRequestStatusCanceled {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.AlreadyCanceled)
		return conflictErrors
	} else if maintenanceRequest.Status != enum.MaintenanceRequestStatusPending {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.NotActive)
		return conflictErrors
	}

	maintenanceRequest.Status = enum.MaintenanceRequestStatusCanceled

	err = maintenanceService.maintenanceRepository.UpdateMaintenanceRequest(maintenanceService.db, maintenanceRequest)
	if err != nil {
		return err
	}
	return nil
}

func (maintenanceService *MaintenanceService) ApproveMaintenanceRecord(maintenanceInfo maintenancedto.CustomerMaintenanceRequest) error {
	maintenanceRequest, err := maintenanceService.getRequest(maintenanceInfo.RequestID)
	if err != nil {
		return err
	}

	_, err = maintenanceService.installationService.ValidatePanelOwnership(maintenanceRequest.PanelID, maintenanceInfo.OwnerID)
	if err != nil {
		return err
	}

	record, err := maintenanceService.getRequestRecord(maintenanceInfo.RequestID)
	if err != nil {
		return err
	}

	if record.IsUserApproved {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRecord, maintenanceService.constants.Tag.AlreadyAccepted)
		return conflictErrors
	}

	record.IsUserApproved = true

	if err = maintenanceService.maintenanceRepository.UpdateMaintenanceRecord(maintenanceService.db, record); err != nil {
		return err
	}
	return nil
}

func (maintenanceService *MaintenanceService) GetCorporationMaintenanceRequests(listInfo maintenancedto.CorporationMaintenanceListRequest) ([]maintenancedto.CorporationMaintenanceListResponse, error) {
	err := maintenanceService.corporationService.CheckApplicantAccess(listInfo.CorporationID, listInfo.OperatorID)
	if err != nil {
		return nil, err
	}

	allowedStatus := maintenanceService.mapStatusForRole(listInfo.Status, enum.AgentTypeCorporation)

	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	maintenanceRequests, err := maintenanceService.maintenanceRepository.FindCorporationRequestsByStatus(maintenanceService.db, listInfo.CorporationID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]maintenancedto.CorporationMaintenanceListResponse, len(maintenanceRequests))

	for i, maintenanceRequest := range maintenanceRequests {
		panelInfoRequest := installationdto.CorporationPanelRequest{
			CorporationID:  listInfo.CorporationID,
			OperatorID:     listInfo.OperatorID,
			InstallationID: maintenanceRequest.PanelID,
		}
		panel, err := maintenanceService.installationService.GetCorporationPanel(panelInfoRequest)
		if err != nil {
			return nil, err
		}

		response[i] = maintenancedto.CorporationMaintenanceListResponse{
			ID:                   maintenanceRequest.ID,
			CreatedAt:            maintenanceRequest.CreatedAt,
			Panel:                panel,
			Subject:              maintenanceRequest.Subject,
			Description:          maintenanceRequest.Description,
			UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
			Status:               maintenanceRequest.Status.String(),
			IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		}
	}
	return response, nil
}

func (maintenanceService *MaintenanceService) GetCorporationMaintenanceRequest(maintenanceInfo maintenancedto.CorporationMaintenanceRequest) (maintenancedto.CorporationMaintenanceResponse, error) {
	if err := maintenanceService.corporationService.CheckApplicantAccess(maintenanceInfo.CorporationID, maintenanceInfo.OperatorID); err != nil {
		return maintenancedto.CorporationMaintenanceResponse{}, err
	}

	maintenanceRequest, err := maintenanceService.getCorporationRequest(maintenanceInfo.RequestID, maintenanceInfo.CorporationID)
	if err != nil {
		return maintenancedto.CorporationMaintenanceResponse{}, err
	}

	panelInfoRequest := installationdto.CorporationPanelRequest{
		CorporationID:  maintenanceInfo.CorporationID,
		OperatorID:     maintenanceInfo.OperatorID,
		InstallationID: maintenanceRequest.PanelID,
	}
	panel, err := maintenanceService.installationService.GetCorporationPanel(panelInfoRequest)
	if err != nil {
		return maintenancedto.CorporationMaintenanceResponse{}, err
	}

	record, err := maintenanceService.getCorporationMaintenanceRecord(maintenanceInfo.RequestID, maintenanceRequest.PanelID)
	if err != nil {
		return maintenancedto.CorporationMaintenanceResponse{}, err
	}

	response := maintenancedto.CorporationMaintenanceResponse{
		ID:                   maintenanceRequest.ID,
		CreatedAt:            maintenanceRequest.CreatedAt,
		Panel:                panel,
		Subject:              maintenanceRequest.Subject,
		Description:          maintenanceRequest.Description,
		UrgencyLevel:         maintenanceRequest.UrgencyLevel.String(),
		Status:               maintenanceRequest.Status.String(),
		IsGuaranteeRequested: maintenanceRequest.IsGuaranteeRequested,
		Record:               record,
	}
	return response, nil
}

func (maintenanceService *MaintenanceService) getCorporationMaintenanceRecord(requestID, panelID uint) (maintenancedto.CorporationMaintenanceRecordResponse, error) {
	record, err := maintenanceService.getRequestRecord(requestID)
	if err != nil {
		return maintenancedto.CorporationMaintenanceRecordResponse{}, err
	}

	operator, err := maintenanceService.userService.GetUserCredential(record.OperatorID)
	if err != nil {
		return maintenancedto.CorporationMaintenanceRecordResponse{}, err
	}

	violation, err := maintenanceService.guaranteeService.GetCorporationPanelGuaranteeViolation(panelID)
	if err != nil {
		return maintenancedto.CorporationMaintenanceRecordResponse{}, err
	}

	recordResponse := maintenancedto.CorporationMaintenanceRecordResponse{
		ID:                 record.ID,
		CreatedAt:          record.CreatedAt,
		Operator:           operator,
		Title:              record.Title,
		Details:            record.Details,
		IsUserApproved:     record.IsUserApproved,
		GuaranteeViolation: violation,
	}
	return recordResponse, nil
}

// TODO: CHECKED COULD BE BETTER add timer
func (maintenanceService *MaintenanceService) AcceptMaintenanceRequest(maintenanceInfo maintenancedto.CorporationMaintenanceRequest) error {
	if err := maintenanceService.corporationService.CheckApplicantAccess(maintenanceInfo.CorporationID, maintenanceInfo.OperatorID); err != nil {
		return err
	}

	maintenanceRequest, err := maintenanceService.getCorporationRequest(maintenanceInfo.RequestID, maintenanceInfo.CorporationID)
	if err != nil {
		return err
	}

	var conflictErrors exception.ConflictErrors
	if maintenanceRequest.Status == enum.MaintenanceRequestStatusAccepted {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.AlreadyAccepted)
		return conflictErrors
	} else if maintenanceRequest.Status == enum.MaintenanceRequestStatusRejected {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.AlreadyRejected)
		return conflictErrors
	}

	maintenanceRequest.Status = enum.MaintenanceRequestStatusAccepted

	if err := maintenanceService.maintenanceRepository.UpdateMaintenanceRequest(maintenanceService.db, maintenanceRequest); err != nil {
		return err
	}
	return nil
}

// TODO: CHECKED COULD BE BETTER add reason
func (maintenanceService *MaintenanceService) RejectMaintenanceRequest(maintenanceInfo maintenancedto.CorporationMaintenanceRequest) error {
	if err := maintenanceService.corporationService.CheckApplicantAccess(maintenanceInfo.CorporationID, maintenanceInfo.OperatorID); err != nil {
		return err
	}

	maintenanceRequest, err := maintenanceService.getCorporationRequest(maintenanceInfo.RequestID, maintenanceInfo.CorporationID)
	if err != nil {
		return err
	}

	var conflictErrors exception.ConflictErrors
	if maintenanceRequest.Status == enum.MaintenanceRequestStatusRejected {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.AlreadyRejected)
		return conflictErrors
	} else if maintenanceRequest.Status == enum.MaintenanceRequestStatusAccepted {
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRequest, maintenanceService.constants.Tag.AlreadyAccepted)
		return conflictErrors
	}

	maintenanceRequest.Status = enum.MaintenanceRequestStatusRejected

	if err := maintenanceService.maintenanceRepository.UpdateMaintenanceRequest(maintenanceService.db, maintenanceRequest); err != nil {
		return err
	}
	return nil
}

func (maintenanceService *MaintenanceService) CreateMaintenanceRecord(recordInfo maintenancedto.CreateMaintenanceRecordRequest) error {
	if err := maintenanceService.corporationService.CheckApplicantAccess(recordInfo.CorporationID, recordInfo.OperatorID); err != nil {
		return err
	}

	maintenanceRequest, err := maintenanceService.getCorporationRequest(recordInfo.RequestID, recordInfo.CorporationID)
	if err != nil {
		return err
	}

	record, err := maintenanceService.maintenanceRepository.FindRecordByRequestID(maintenanceService.db, recordInfo.RequestID)
	if err != nil {
		return err
	}
	if record != nil {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRecord, maintenanceService.constants.Tag.AlreadyExist)
		return conflictErrors
	}

	err = maintenanceService.db.WithTransaction(func(tx database.Database) error {
		var guaranteeViolationID *uint = nil
		if recordInfo.GuaranteeViolation != nil {
			recordInfo.GuaranteeViolation.PanelID = maintenanceRequest.PanelID
			request := installationdto.CreateViolatePanelGuaranteeRequest{
				CorporationID:      recordInfo.CorporationID,
				OperatorID:         recordInfo.OperatorID,
				PanelID:            maintenanceRequest.PanelID,
				GuaranteeViolation: *recordInfo.GuaranteeViolation,
			}
			temp, err := maintenanceService.installationService.ViolatePanelGuaranteeStatus(request)
			if err != nil {
				return err
			}
			guaranteeViolationID = &temp
		}

		record = &entity.MaintenanceRecord{
			OperatorID:           recordInfo.OperatorID,
			RequestID:            recordInfo.RequestID,
			IsUserApproved:       false,
			Title:                recordInfo.Title,
			Details:              recordInfo.Details,
			GuaranteeViolationID: guaranteeViolationID,
		}
		if err := maintenanceService.maintenanceRepository.CreateMaintenanceRecord(tx, record); err != nil {
			return err
		}

		maintenanceRequest.Status = enum.MaintenanceRequestStatusCompleted
		if err := maintenanceService.maintenanceRepository.UpdateMaintenanceRequest(tx, maintenanceRequest); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (maintenanceService *MaintenanceService) UpdateMaintenanceRecord(recordInfo maintenancedto.UpdateMaintenanceRecordRequest) error {
	if err := maintenanceService.corporationService.CheckApplicantAccess(recordInfo.CorporationID, recordInfo.OperatorID); err != nil {
		return err
	}

	maintenanceRequest, err := maintenanceService.getCorporationRequest(recordInfo.RequestID, recordInfo.CorporationID)
	if err != nil {
		return err
	}

	maintenanceRecord, err := maintenanceService.getRequestRecord(recordInfo.RequestID)
	if err != nil {
		return err
	}

	if maintenanceRecord.IsUserApproved {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(maintenanceService.constants.Field.MaintenanceRecord, maintenanceService.constants.Tag.NotActive)
		return conflictErrors
	}

	if recordInfo.Title != nil {
		maintenanceRecord.Title = *recordInfo.Title
	}
	if recordInfo.Details != nil {
		maintenanceRecord.Details = *recordInfo.Details
	}

	err = maintenanceService.db.WithTransaction(func(tx database.Database) error {
		if recordInfo.GuaranteeViolation != nil {
			recordInfo.GuaranteeViolation.PanelID = maintenanceRequest.PanelID
			if err := maintenanceService.guaranteeService.UpdateGuaranteeViolation(*recordInfo.GuaranteeViolation); err != nil {
				return err
			}
		}

		if err := maintenanceService.maintenanceRepository.UpdateMaintenanceRecord(tx, maintenanceRecord); err != nil {
			return err
		}
		return nil
	})
	return err
}
