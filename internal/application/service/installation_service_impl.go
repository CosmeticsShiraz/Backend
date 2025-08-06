package service

import (
	"time"

	"github.com/CosmeticsShiraz/Backend/bootstrap"
	chatdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/chat"
	guaranteedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/guarantee"
	installationdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/installation"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/domain/exception"
	"github.com/CosmeticsShiraz/Backend/internal/domain/repository/postgres"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
	postgresImpl "github.com/CosmeticsShiraz/Backend/internal/infrastructure/repository/postgres"
)

type InstallationService struct {
	constants              *bootstrap.Constants
	addressService         usecase.AddressService
	userService            usecase.UserService
	corporationService     usecase.CorporationService
	guaranteeService       usecase.GuaranteeService
	chatService            usecase.ChatService
	installationRepository postgres.InstallationRepository
	db                     database.Database
}

type InstallationServiceDeps struct {
	Constants              *bootstrap.Constants
	AddressService         usecase.AddressService
	UserService            usecase.UserService
	CorporationService     usecase.CorporationService
	GuaranteeService       usecase.GuaranteeService
	ChatService            usecase.ChatService
	InstallationRepository postgres.InstallationRepository
	DB                     database.Database
}

func NewInstallationService(deps InstallationServiceDeps) *InstallationService {
	return &InstallationService{
		constants:              deps.Constants,
		addressService:         deps.AddressService,
		userService:            deps.UserService,
		corporationService:     deps.CorporationService,
		guaranteeService:       deps.GuaranteeService,
		chatService:            deps.ChatService,
		installationRepository: deps.InstallationRepository,
		db:                     deps.DB,
	}
}

func (installationService *InstallationService) getOwnerRequest(requestID, ownerID uint) (*entity.InstallationRequest, error) {
	request, err := installationService.installationRepository.FindRequestByOwner(installationService.db, requestID, ownerID)
	if err != nil {
		return nil, err
	}
	if request == nil {
		return nil, exception.NotFoundError{Item: installationService.constants.Field.InstallationRequest}
	}
	return request, nil
}

func (installationService *InstallationService) getRequest(requestID uint) (*entity.InstallationRequest, error) {
	request, err := installationService.installationRepository.FindRequestByID(installationService.db, requestID)
	if err != nil {
		return nil, err
	}
	if request == nil {
		return nil, exception.NotFoundError{Item: installationService.constants.Field.InstallationRequest}
	}
	return request, nil
}

func (installationService *InstallationService) getCorporationPanel(panelID, corporationID uint) (*entity.Panel, error) {
	panel, err := installationService.installationRepository.FindCorporationPanel(installationService.db, panelID, corporationID)
	if err != nil {
		return nil, err
	}
	if panel == nil {
		return nil, exception.NotFoundError{Item: installationService.constants.Field.Panel}
	}
	return panel, nil
}

func (installationService *InstallationService) getCustomerPanel(panelID, ownerID uint) (*entity.Panel, error) {
	panel, err := installationService.installationRepository.FindCustomerPanel(installationService.db, panelID, ownerID)
	if err != nil {
		return nil, err
	}
	if panel == nil {
		return nil, exception.NotFoundError{Item: installationService.constants.Field.Panel}
	}
	return panel, nil
}

func (installationService *InstallationService) getPanel(panelID uint) (*entity.Panel, error) {
	panel, err := installationService.installationRepository.FindPanelByID(installationService.db, panelID)
	if err != nil {
		return nil, err
	}
	if panel == nil {
		return nil, exception.NotFoundError{Item: installationService.constants.Field.Panel}
	}
	return panel, nil
}

func (installationService *InstallationService) GetRequestStatuses() []installationdto.EnumStatusResponse {
	statuses := enum.GetAllInstallationRequestStatuses()
	response := make([]installationdto.EnumStatusResponse, len(statuses))
	for i, status := range statuses {
		response[i] = installationdto.EnumStatusResponse{
			ID:   uint(status),
			Name: status.String(),
		}
	}
	return response
}

func (installationService *InstallationService) GetBuildingTypes() []installationdto.EnumStatusResponse {
	types := enum.GetAllBuildingTypes()
	response := make([]installationdto.EnumStatusResponse, len(types))
	for i, buildingType := range types {
		response[i] = installationdto.EnumStatusResponse{
			ID:   uint(buildingType),
			Name: buildingType.String(),
		}
	}
	return response
}

func (installationService *InstallationService) GetPanelStatuses() []installationdto.EnumStatusResponse {
	statuses := enum.GetAllPanelStatuses()
	response := make([]installationdto.EnumStatusResponse, len(statuses))
	for i, status := range statuses {
		response[i] = installationdto.EnumStatusResponse{
			ID:   uint(status),
			Name: status.String(),
		}
	}
	return response
}

func (installationService *InstallationService) ValidateRequestOwnership(requestID, ownerID uint) (installationdto.PublicRequestDetailsResponse, error) {
	request, err := installationService.getOwnerRequest(requestID, ownerID)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}

	customer, err := installationService.userService.GetUserCredential(request.OwnerID)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}
	address, err := installationService.addressService.GetAddress(request.ID, installationService.constants.AddressOwners.InstallationRequest)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}

	return installationdto.PublicRequestDetailsResponse{
		ID:           request.ID,
		Name:         request.Name,
		Status:       request.Status.String(),
		PowerRequest: request.PowerRequest,
		Description:  request.Description,
		BuildingType: request.BuildingType.String(),
		Area:         request.Area,
		MaxCost:      request.MaxCost,
		Customer:     customer,
		Address:      address,
	}, nil
}

func (installationService *InstallationService) CreateInstallationRequest(request installationdto.NewInstallationRequest) error {
	// compare installed panels names to new request name
	if err := installationService.userService.IsUserActive(request.OwnerID); err != nil {
		return err
	}

	allowedStatus := []enum.InstallationRequestStatus{enum.InstallationRequestStatusActive}
	existingRequest, err := installationService.installationRepository.FindOwnerRequestByName(installationService.db, request.OwnerID, allowedStatus, request.Name)
	if err != nil {
		return err
	}
	if existingRequest != nil {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(installationService.constants.Field.Name, installationService.constants.Tag.AlreadyRegistered)
		return conflictErrors
	}
	inProgressReqs, err := installationService.installationRepository.FindOwnerRequests(installationService.db, request.OwnerID, allowedStatus)
	if err != nil {
		return err
	}
	if len(inProgressReqs) >= 5 {
		rateLimitError := exception.NewConcurrentInstallLimitError("", 5, nil)
		return rateLimitError
	}

	installationRequest := &entity.InstallationRequest{
		Name:         request.Name,
		Status:       enum.InstallationRequestStatusActive,
		Area:         request.Area,
		PowerRequest: request.Power,
		MaxCost:      request.MaxCost,
		BuildingType: enum.BuildingType(request.BuildingType),
		OwnerID:      request.OwnerID,
		Description:  request.Description,
		Address: entity.Address{
			ProvinceID:    request.Address.ProvinceID,
			CityID:        request.Address.CityID,
			StreetAddress: request.Address.StreetAddress,
			PostalCode:    request.Address.PostalCode,
			HouseNumber:   request.Address.HouseNumber,
			Unit:          request.Address.Unit,
		},
	}
	err = installationService.db.WithTransaction(func(tx database.Database) error {
		err = installationService.installationRepository.CreateRequest(installationService.db, installationRequest)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (installationService *InstallationService) GetOwnerInstallationRequests(request installationdto.CustomerRequestsListRequest) ([]installationdto.AnonymousRequestsResponse, error) {
	status := enum.InstallationRequestStatus(request.Status)
	allowedStatus := []enum.InstallationRequestStatus{status}
	if status == enum.InstallationRequestStatusAll {
		allowedStatus = enum.GetAllInstallationRequestStatuses()
	}

	paginationModifier := postgresImpl.NewPaginationModifier(request.Limit, request.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	requests, err := installationService.installationRepository.FindOwnerRequests(
		installationService.db, request.OwnerID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.AnonymousRequestsResponse, len(requests))

	for i, request := range requests {
		address, err := installationService.addressService.GetAddress(request.ID, installationService.constants.AddressOwners.InstallationRequest)
		if err != nil {
			return nil, err
		}
		response[i] = installationdto.AnonymousRequestsResponse{
			ID:           request.ID,
			Name:         request.Name,
			CreatedTime:  request.CreatedAt,
			Status:       request.Status.String(),
			PowerRequest: request.PowerRequest,
			MaxCost:      request.MaxCost,
			BuildingType: request.BuildingType.String(),
			Address:      address,
		}
	}
	return response, nil
}

func (installationService *InstallationService) GetOwnerInstallationRequest(request installationdto.GetOwnerRequest) (installationdto.AnonymousRequestsResponse, error) {
	installationRequest, err := installationService.getOwnerRequest(request.InstallationID, request.OwnerID)
	if err != nil {
		return installationdto.AnonymousRequestsResponse{}, err
	}

	address, err := installationService.addressService.GetAddress(request.InstallationID, installationService.constants.AddressOwners.InstallationRequest)
	if err != nil {
		return installationdto.AnonymousRequestsResponse{}, err
	}
	response := installationdto.AnonymousRequestsResponse{
		ID:           installationRequest.ID,
		Name:         installationRequest.Name,
		CreatedTime:  installationRequest.CreatedAt,
		Status:       installationRequest.Status.String(),
		PowerRequest: installationRequest.PowerRequest,
		MaxCost:      installationRequest.MaxCost,
		BuildingType: installationRequest.BuildingType.String(),
		Address:      address,
	}
	return response, nil
}

func (installationService *InstallationService) ChangeInstallationRequestStatus(request installationdto.ChangeRequestStatusRequest) error {
	installationRequest, err := installationService.getOwnerRequest(request.RequestID, request.OwnerID)
	if err != nil {
		return err
	}

	if installationRequest.Status == enum.InstallationRequestStatusCancelled {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(installationService.constants.Field.InstallationRequest, installationService.constants.Tag.AlreadyCanceled)
		return conflictErrors
	}

	installationRequest.Status = request.Status
	err = installationService.installationRepository.UpdateRequest(installationService.db, installationRequest)
	if err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) GetAnonymousInstallationRequests(request installationdto.CorporationPanelListRequest) ([]installationdto.AnonymousRequestsResponse, error) {
	if err := installationService.corporationService.CheckApplicantAccess(request.CorporationID, request.OperatorID); err != nil {
		return nil, err
	}

	allowedStatus := []enum.InstallationRequestStatus{enum.InstallationRequestStatusActive}

	paginationModifier := postgresImpl.NewPaginationModifier(request.Limit, request.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	installationRequests, err := installationService.installationRepository.FindRequestsByStatus(installationService.db, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.AnonymousRequestsResponse, len(installationRequests))

	for i, installationRequest := range installationRequests {
		address, err := installationService.addressService.GetAddress(installationRequest.ID, installationService.constants.AddressOwners.InstallationRequest)
		if err != nil {
			return nil, err
		}
		response[i] = installationdto.AnonymousRequestsResponse{
			ID:           installationRequest.ID,
			Name:         installationRequest.Name,
			CreatedTime:  installationRequest.CreatedAt,
			Status:       installationRequest.Status.String(),
			PowerRequest: installationRequest.PowerRequest,
			MaxCost:      installationRequest.MaxCost,
			BuildingType: installationRequest.BuildingType.String(),
			Address:      address,
		}
	}
	return response, nil
}

func (installationService *InstallationService) GetAnonymousInstallationRequest(request installationdto.CorporationPanelRequest) (installationdto.AnonymousRequestsResponse, error) {
	if err := installationService.corporationService.CheckApplicantAccess(request.CorporationID, request.OperatorID); err != nil {
		return installationdto.AnonymousRequestsResponse{}, err
	}

	installationRequest, err := installationService.getRequest(request.InstallationID)
	if err != nil {
		return installationdto.AnonymousRequestsResponse{}, err
	}
	if installationRequest.Status != enum.InstallationRequestStatusActive {
		return installationdto.AnonymousRequestsResponse{}, exception.NotFoundError{Item: installationService.constants.Field.InstallationRequest}
	}

	address, err := installationService.addressService.GetAddress(installationRequest.ID, installationService.constants.AddressOwners.InstallationRequest)
	if err != nil {
		return installationdto.AnonymousRequestsResponse{}, err
	}

	response := installationdto.AnonymousRequestsResponse{
		ID:           installationRequest.ID,
		Name:         installationRequest.Name,
		CreatedTime:  installationRequest.CreatedAt,
		Status:       installationRequest.Status.String(),
		PowerRequest: installationRequest.PowerRequest,
		MaxCost:      installationRequest.MaxCost,
		BuildingType: installationRequest.BuildingType.String(),
		Address:      address,
	}
	return response, nil
}

func (installationService *InstallationService) GetPublicInstallationRequest(requestID uint) (installationdto.PublicRequestDetailsResponse, error) {
	installationRequest, err := installationService.getRequest(requestID)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}

	customer, err := installationService.userService.GetUserCredential(installationRequest.OwnerID)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}

	address, err := installationService.addressService.GetAddress(installationRequest.ID, installationService.constants.AddressOwners.InstallationRequest)
	if err != nil {
		return installationdto.PublicRequestDetailsResponse{}, err
	}

	response := installationdto.PublicRequestDetailsResponse{
		ID:           installationRequest.ID,
		Name:         installationRequest.Name,
		Status:       installationRequest.Status.String(),
		PowerRequest: installationRequest.PowerRequest,
		Description:  installationRequest.Description,
		BuildingType: installationRequest.BuildingType.String(),
		Area:         installationRequest.Area,
		MaxCost:      installationRequest.MaxCost,
		Customer:     customer,
		Address:      address,
	}
	return response, nil
}

func (installationService *InstallationService) GetInstallationRequestsByAdmin(request installationdto.AdminInstallationListRequest) ([]installationdto.PublicRequestDetailsResponse, error) {
	allowedStatuses := []enum.InstallationRequestStatus{enum.InstallationRequestStatus(request.Status)}
	if enum.InstallationRequestStatus(request.Status) == enum.InstallationRequestStatusAll {
		allowedStatuses = enum.GetAllInstallationRequestStatuses()
	}

	paginationModifier := postgresImpl.NewPaginationModifier(request.Limit, request.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	installationRequests, err := installationService.installationRepository.FindRequestsByStatus(installationService.db, allowedStatuses, sortingModifier, paginationModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.PublicRequestDetailsResponse, len(installationRequests))

	for i, installationRequest := range installationRequests {
		customer, err := installationService.userService.GetUserCredential(installationRequest.OwnerID)
		if err != nil {
			return nil, err
		}

		address, err := installationService.addressService.GetAddress(installationRequest.ID, installationService.constants.AddressOwners.InstallationRequest)
		if err != nil {
			return nil, err
		}

		response[i] = installationdto.PublicRequestDetailsResponse{
			ID:           installationRequest.ID,
			Name:         installationRequest.Name,
			Status:       installationRequest.Status.String(),
			PowerRequest: installationRequest.PowerRequest,
			Description:  installationRequest.Description,
			BuildingType: installationRequest.BuildingType.String(),
			Area:         installationRequest.Area,
			MaxCost:      installationRequest.MaxCost,
			Customer:     customer,
			Address:      address,
		}
	}

	return response, nil
}

func (installationService *InstallationService) CompleteInstallationRequest(request installationdto.CompleteBidInstallationRequest) error {
	if err := installationService.corporationService.CheckApplicantAccess(request.CorporationID, request.OperatorID); err != nil {
		return err
	}

	if err := installationService.userService.IsUserActive(request.OperatorID); err != nil {
		return err
	}

	panel, err := installationService.getCorporationPanel(request.PanelID, request.CorporationID)
	if err != nil {
		return err
	}

	panel.Tilt = request.Tilt
	panel.Azimuth = request.Azimuth
	panel.TotalNumberOfModules = request.NumberOfModules

	if err := installationService.installationRepository.UpdatePanel(installationService.db, panel); err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) UpdateInstallationRequestByAdmin(newRequest installationdto.UpdateInstallationRequest) error {
	installationRequest, err := installationService.getRequest(newRequest.RequestID)
	if err != nil {
		return err
	}

	if newRequest.Name != nil {
		installationRequest.Name = *newRequest.Name
	}
	if newRequest.Area != nil {
		installationRequest.Area = *newRequest.Area
	}
	if newRequest.Power != nil {
		installationRequest.PowerRequest = *newRequest.Power
	}
	if newRequest.MaxCost != nil {
		installationRequest.MaxCost = *newRequest.MaxCost
	}
	if newRequest.BuildingType != nil {
		installationRequest.BuildingType = enum.BuildingType(*newRequest.BuildingType)
	}
	if newRequest.Status != nil {
		installationRequest.Status = enum.InstallationRequestStatus(*newRequest.Status)
	}
	if newRequest.Description != nil {
		installationRequest.Description = *newRequest.Description
	}

	if err := installationService.installationRepository.UpdateRequest(installationService.db, installationRequest); err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) DeleteInstallationRequest(requestID uint) error {
	installationRequest, err := installationService.getRequest(requestID)
	if err != nil {
		return err
	}

	if err := installationService.installationRepository.DeleteRequest(installationService.db, installationRequest); err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) ValidatePanelOwnership(panelID, userID uint) (installationdto.AdminPanelResponse, error) {
	panel, err := installationService.installationRepository.FindPanelByOwner(installationService.db, panelID, userID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	if panel == nil {
		return installationdto.AdminPanelResponse{}, exception.NotFoundError{Item: installationService.constants.Field.Panel}
	}
	operator, err := installationService.userService.GetUserCredential(panel.OperatorID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	customer, err := installationService.userService.GetUserCredential(panel.CustomerID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	corporation, err := installationService.corporationService.GetCorporationCredentials(panel.CorporationID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	var guarantee guaranteedto.GuaranteeResponse
	if panel.GuaranteeID != nil {
		guarantee, err = installationService.guaranteeService.GetGuarantee(*panel.GuaranteeID)
		if err != nil {
			return installationdto.AdminPanelResponse{}, err
		}
	}
	response := installationdto.AdminPanelResponse{
		ID:                   panel.ID,
		Name:                 panel.Name,
		Status:               panel.Status.String(),
		BuildingType:         panel.BuildingType.String(),
		Area:                 panel.Area,
		Power:                panel.Power,
		Tilt:                 panel.Tilt,
		Azimuth:              panel.Azimuth,
		TotalNumberOfModules: panel.TotalNumberOfModules,
		GuaranteeStatus:      panel.GuaranteeStatus.String(),
		Operator:             operator,
		Customer:             customer,
		Address:              address,
		Guarantee:            guarantee,
		Corporation:          corporation,
	}
	return response, nil
}

func (installationService *InstallationService) ValidatePanelGuarantee(panelID uint) error {
	panel, err := installationService.getPanel(panelID)
	if err != nil {
		return err
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusActive {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(installationService.constants.Field.Guarantee, installationService.constants.Tag.NotActive)
		return conflictErrors
	}
	return nil
}

// TODO: nor done remain for the bid/bidID/request && maybe remove:FindPanelByNameAndCustomerID
func (installationService *InstallationService) AddPanel(panelInfo installationdto.AddPanelRequest) error {
	if err := installationService.corporationService.ISCorporationApproved(panelInfo.CorporationID); err != nil {
		return err
	}

	if err := installationService.corporationService.CheckApplicantAccess(panelInfo.CorporationID, panelInfo.OperatorID); err != nil {
		return err
	}

	if err := installationService.userService.IsUserActive(panelInfo.OperatorID); err != nil {
		return err
	}

	customer, err := installationService.userService.FindActiveUserByPhone(panelInfo.CustomerPhone)
	if err != nil {
		return err
	}

	panel, err := installationService.installationRepository.FindPanelByNameAndCustomerID(installationService.db, panelInfo.Name, customer.ID)
	if err != nil {
		return err
	}
	if panel != nil {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(installationService.constants.Field.Name, installationService.constants.Tag.AlreadyExist)
		return conflictErrors
	}

	panelGuaranteeStatus := enum.PanelGuaranteeStatusEmpty
	if panelInfo.GuaranteeID != nil {
		if err := installationService.guaranteeService.ValidateActiveGuaranteeOwnerShip(*panelInfo.GuaranteeID, panelInfo.CorporationID); err != nil {
			return err
		}
		panelGuaranteeStatus = enum.PanelGuaranteeStatusActive
	}

	panel = &entity.Panel{
		Name:                 panelInfo.Name,
		Status:               panelInfo.Status,
		BuildingType:         enum.BuildingType(panelInfo.BuildingType),
		Power:                panelInfo.Power,
		Area:                 panelInfo.Area,
		Tilt:                 panelInfo.Tilt,
		Azimuth:              panelInfo.Azimuth,
		TotalNumberOfModules: panelInfo.TotalNumberOfModules,
		CorporationID:        panelInfo.CorporationID,
		OperatorID:           panelInfo.OperatorID,
		CustomerID:           customer.ID,
		GuaranteeStatus:      panelGuaranteeStatus,
		GuaranteeID:          panelInfo.GuaranteeID,
		Address: entity.Address{
			ProvinceID:    panelInfo.Address.ProvinceID,
			CityID:        panelInfo.Address.CityID,
			StreetAddress: panelInfo.Address.StreetAddress,
			PostalCode:    panelInfo.Address.PostalCode,
			HouseNumber:   panelInfo.Address.HouseNumber,
			Unit:          panelInfo.Address.Unit,
		},
	}
	err = installationService.db.WithTransaction(func(tx database.Database) error {
		if err = installationService.installationRepository.CreatePanel(installationService.db, panel); err != nil {
			return err
		}
		request := chatdto.CreateOrGetUserRoomRequest{
			CorporationID: panel.CorporationID,
			UserID:        customer.ID,
		}
		if _, err := installationService.chatService.CreateOrGetRoom(request); err != nil {
			return err
		}
		return nil
	})

	return err
}

func (installationService *InstallationService) GetCorporationPanels(listInfo installationdto.CorporationPanelListRequest) ([]installationdto.CorporationPanelListResponse, error) {
	if err := installationService.corporationService.CheckApplicantAccess(listInfo.CorporationID, listInfo.OperatorID); err != nil {
		return nil, err
	}

	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	allowedStatus := []enum.PanelStatus{enum.PanelStatus(listInfo.Status)}
	if enum.PanelStatus(listInfo.Status) == enum.PanelStatusAll {
		allowedStatus = enum.GetAllPanelStatuses()
	}

	panels, err := installationService.installationRepository.FindCorporationPanels(installationService.db, listInfo.CorporationID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.CorporationPanelListResponse, len(panels))

	for i, panel := range panels {
		address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
		if err != nil {
			return nil, err
		}
		customer, err := installationService.userService.GetUserCredential(panel.CustomerID)
		if err != nil {
			return nil, err
		}
		operator, err := installationService.userService.GetUserCredential(panel.OperatorID)
		if err != nil {
			return nil, err
		}
		response[i] = installationdto.CorporationPanelListResponse{
			ID:                   panel.ID,
			Name:                 panel.Name,
			Status:               panel.Status.String(),
			BuildingType:         panel.BuildingType.String(),
			Area:                 panel.Area,
			Power:                panel.Power,
			Tilt:                 panel.Tilt,
			Azimuth:              panel.Azimuth,
			TotalNumberOfModules: panel.TotalNumberOfModules,
			GuaranteeStatus:      panel.GuaranteeStatus.String(),
			Operator:             operator,
			Customer:             customer,
			Address:              address,
		}
	}
	return response, nil
}

func (installationService *InstallationService) GetCorporationPanel(request installationdto.CorporationPanelRequest) (installationdto.CorporationPanelResponse, error) {
	if err := installationService.corporationService.CheckApplicantAccess(request.CorporationID, request.OperatorID); err != nil {
		return installationdto.CorporationPanelResponse{}, err
	}

	panel, err := installationService.getCorporationPanel(request.InstallationID, request.CorporationID)
	if err != nil {
		return installationdto.CorporationPanelResponse{}, err
	}

	address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
	if err != nil {
		return installationdto.CorporationPanelResponse{}, err
	}
	customer, err := installationService.userService.GetUserCredential(panel.CustomerID)
	if err != nil {
		return installationdto.CorporationPanelResponse{}, err
	}
	operator, err := installationService.userService.GetUserCredential(panel.OperatorID)
	if err != nil {
		return installationdto.CorporationPanelResponse{}, err
	}

	var guarantee guaranteedto.GuaranteeResponse
	if panel.GuaranteeID != nil {
		guarantee, err = installationService.guaranteeService.GetGuarantee(*panel.GuaranteeID)
		if err != nil {
			return installationdto.CorporationPanelResponse{}, err
		}
	}

	response := installationdto.CorporationPanelResponse{
		ID:                   panel.ID,
		Name:                 panel.Name,
		Status:               panel.Status.String(),
		BuildingType:         panel.BuildingType.String(),
		Area:                 panel.Area,
		Power:                panel.Power,
		Tilt:                 panel.Tilt,
		Azimuth:              panel.Azimuth,
		TotalNumberOfModules: panel.TotalNumberOfModules,
		GuaranteeStatus:      panel.GuaranteeStatus.String(),
		Operator:             operator,
		Customer:             customer,
		Address:              address,
		Guarantee:            guarantee,
	}
	return response, nil
}

func (installationService *InstallationService) GetCustomerPanels(listInfo installationdto.CustomerPanelListRequest) ([]installationdto.CustomerPanelListResponse, error) {
	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	allowedStatus := []enum.PanelStatus{enum.PanelStatus(listInfo.Status)}
	if enum.PanelStatus(listInfo.Status) == enum.PanelStatusAll {
		allowedStatus = enum.GetAllPanelStatuses()
	}

	panels, err := installationService.installationRepository.FindCustomerPanels(installationService.db, listInfo.OwnerID, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.CustomerPanelListResponse, len(panels))

	for i, panel := range panels {
		address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
		if err != nil {
			return nil, err
		}
		corporation, err := installationService.corporationService.GetCorporationCredentials(panel.CorporationID)
		if err != nil {
			return nil, err
		}

		response[i] = installationdto.CustomerPanelListResponse{
			ID:                   panel.ID,
			Name:                 panel.Name,
			Status:               panel.Status.String(),
			BuildingType:         panel.BuildingType.String(),
			Area:                 panel.Area,
			Power:                panel.Power,
			Tilt:                 panel.Tilt,
			Azimuth:              panel.Azimuth,
			TotalNumberOfModules: panel.TotalNumberOfModules,
			GuaranteeStatus:      panel.GuaranteeStatus.String(),
			Corporation:          corporation,
			Address:              address,
		}
	}
	return response, nil
}

func (installationService *InstallationService) GetCustomerPanel(panelInfo installationdto.GetOwnerRequest) (installationdto.CustomerPanelResponse, error) {
	panel, err := installationService.getCustomerPanel(panelInfo.InstallationID, panelInfo.OwnerID)
	if err != nil {
		return installationdto.CustomerPanelResponse{}, err
	}

	address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
	if err != nil {
		return installationdto.CustomerPanelResponse{}, err
	}
	corporation, err := installationService.corporationService.GetCorporationCredentials(panel.CorporationID)
	if err != nil {
		return installationdto.CustomerPanelResponse{}, err
	}

	var guarantee guaranteedto.GuaranteeResponse
	if panel.GuaranteeID != nil {
		guarantee, err = installationService.guaranteeService.GetGuarantee(*panel.GuaranteeID)
		if err != nil {
			return installationdto.CustomerPanelResponse{}, err
		}
	}

	response := installationdto.CustomerPanelResponse{
		ID:                   panel.ID,
		Name:                 panel.Name,
		Status:               panel.Status.String(),
		BuildingType:         panel.BuildingType.String(),
		Area:                 panel.Area,
		Power:                panel.Power,
		Tilt:                 panel.Tilt,
		Azimuth:              panel.Azimuth,
		TotalNumberOfModules: panel.TotalNumberOfModules,
		GuaranteeStatus:      panel.GuaranteeStatus.String(),
		Corporation:          corporation,
		Address:              address,
		Guarantee:            guarantee,
	}
	return response, nil
}

func (installationService *InstallationService) GetPanelByAdmin(panelID uint) (installationdto.AdminPanelResponse, error) {
	panel, err := installationService.getPanel(panelID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}

	customer, err := installationService.userService.GetUserCredential(panel.CustomerID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	operator, err := installationService.userService.GetUserCredential(panel.OperatorID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	corporation, err := installationService.corporationService.GetCorporationCredentials(panel.CorporationID)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}
	address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
	if err != nil {
		return installationdto.AdminPanelResponse{}, err
	}

	var guarantee guaranteedto.GuaranteeResponse
	if panel.Guarantee != nil {
		guarantee, _ = installationService.guaranteeService.GetGuarantee(*panel.GuaranteeID)
	}

	response := installationdto.AdminPanelResponse{
		ID:                   panelID,
		Name:                 panel.Name,
		Status:               panel.Status.String(),
		BuildingType:         panel.BuildingType.String(),
		Area:                 panel.Area,
		Power:                panel.Power,
		Tilt:                 panel.Tilt,
		Azimuth:              panel.Azimuth,
		TotalNumberOfModules: panel.TotalNumberOfModules,
		GuaranteeStatus:      panel.GuaranteeStatus.String(),
		Operator:             operator,
		Customer:             customer,
		Corporation:          corporation,
		Address:              address,
		Guarantee:            guarantee,
	}
	return response, nil
}

func (installationService *InstallationService) GetPanelStatus() []installationdto.EnumStatusResponse {
	types := enum.GetAllPanelStatuses()
	response := make([]installationdto.EnumStatusResponse, len(types))
	for i, status := range types {
		response[i] = installationdto.EnumStatusResponse{
			ID:   uint(status),
			Name: status.String(),
		}
	}
	return response
}

func (installationService *InstallationService) GetPanelsByAdmin(listInfo installationdto.AdminInstallationListRequest) ([]installationdto.AdminPanelResponse, error) {
	paginationModifier := postgresImpl.NewPaginationModifier(listInfo.Limit, listInfo.Offset)
	sortingModifier := postgresImpl.NewSortingModifier("created_at", true)

	allowedStatus := []enum.PanelStatus{enum.PanelStatus(listInfo.Status)}
	if enum.PanelStatus(listInfo.Status) == enum.PanelStatusAll {
		allowedStatus = enum.GetAllPanelStatuses()
	}

	panels, err := installationService.installationRepository.FindPanelsByStatus(installationService.db, allowedStatus, paginationModifier, sortingModifier)
	if err != nil {
		return nil, err
	}
	response := make([]installationdto.AdminPanelResponse, len(panels))

	for i, panel := range panels {
		customer, err := installationService.userService.GetUserCredential(panel.CustomerID)
		if err != nil {
			return nil, err
		}
		operator, err := installationService.userService.GetUserCredential(panel.OperatorID)
		if err != nil {
			return nil, err
		}
		corporation, err := installationService.corporationService.GetCorporationCredentials(panel.CorporationID)
		if err != nil {
			return nil, err
		}
		address, err := installationService.addressService.GetAddress(panel.ID, installationService.constants.AddressOwners.Panel)
		if err != nil {
			return nil, err
		}

		var guarantee guaranteedto.GuaranteeResponse
		if panel.Guarantee != nil {
			guarantee, _ = installationService.guaranteeService.GetGuarantee(*panel.GuaranteeID)
		}

		response[i] = installationdto.AdminPanelResponse{
			ID:                   panel.ID,
			Name:                 panel.Name,
			Status:               panel.Status.String(),
			BuildingType:         panel.BuildingType.String(),
			Area:                 panel.Area,
			Power:                panel.Power,
			Tilt:                 panel.Tilt,
			Azimuth:              panel.Azimuth,
			TotalNumberOfModules: panel.TotalNumberOfModules,
			GuaranteeStatus:      panel.GuaranteeStatus.String(),
			Operator:             operator,
			Customer:             customer,
			Corporation:          corporation,
			Address:              address,
			Guarantee:            guarantee,
		}
	}

	return response, nil
}

func (installationService *InstallationService) UpdatePanel(request installationdto.UpdatePanelRequest) error {
	panel, err := installationService.getPanel(request.PanelID)
	if err != nil {
		return err
	}

	if request.Name != nil {
		panel.Name = *request.Name
	}
	if request.Status != nil {
		panel.Status = enum.PanelStatus(*request.Status)
	}
	if request.BuildingType != nil {
		panel.BuildingType = enum.BuildingType(*request.BuildingType)
	}
	if request.Area != nil {
		panel.Area = *request.Area
	}
	if request.Power != nil {
		panel.Power = *request.Power
	}
	if request.Tilt != nil {
		panel.Tilt = *request.Tilt
	}
	if request.Azimuth != nil {
		panel.Azimuth = *request.Azimuth
	}
	if request.TotalNumberOfModules != nil {
		panel.TotalNumberOfModules = *request.TotalNumberOfModules
	}

	if err := installationService.installationRepository.UpdatePanel(installationService.db, panel); err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) DeletePanel(panelID uint) error {
	panel, err := installationService.getPanel(panelID)
	if err != nil {
		return err
	}

	if err := installationService.installationRepository.DeletePanel(installationService.db, panel); err != nil {
		return err
	}
	return nil
}

func (installationService *InstallationService) ViolatePanelGuaranteeStatus(request installationdto.CreateViolatePanelGuaranteeRequest) (uint, error) {
	err := installationService.corporationService.CheckApplicantAccess(request.CorporationID, request.OperatorID)
	if err != nil {
		return 0, err
	}

	panel, err := installationService.getCorporationPanel(request.PanelID, request.CorporationID)
	if err != nil {
		return 0, err
	}

	if panel.GuaranteeStatus == enum.PanelGuaranteeStatusEmpty {
		return 0, exception.NotFoundError{Item: installationService.constants.Field.Guarantee}
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusActive {
		var conflictErrors exception.ConflictErrors
		conflictErrors.Add(installationService.constants.Field.Guarantee, installationService.constants.Tag.NotActive)
		return 0, conflictErrors
	}

	var violationID uint
	err = installationService.db.WithTransaction(func(tx database.Database) error {
		var err error
		violationID, err = installationService.guaranteeService.CreateGuaranteeViolation(request.GuaranteeViolation)
		if err != nil {
			return err
		}
		panel.GuaranteeStatus = enum.PanelGuaranteeStatusVoided

		if err := installationService.installationRepository.UpdatePanel(tx, panel); err != nil {
			return err
		}
		return nil
	})

	return violationID, err
}

func (installationService *InstallationService) ClearPanelGuaranteeViolation(violationInfo installationdto.GetCorporationGuaranteeViolationRequest) error {
	err := installationService.corporationService.CheckApplicantAccess(violationInfo.CorporationID, violationInfo.OperatorID)
	if err != nil {
		return err
	}

	panel, err := installationService.getCorporationPanel(violationInfo.PanelID, violationInfo.CorporationID)
	if err != nil {
		return err
	}

	if panel.GuaranteeStatus == enum.PanelGuaranteeStatusEmpty {
		return exception.NotFoundError{Item: installationService.constants.Field.Guarantee}
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusVoided {
		return exception.NotFoundError{Item: installationService.constants.Field.GuaranteeViolation}
	}

	err = installationService.db.WithTransaction(func(tx database.Database) error {
		if err := installationService.guaranteeService.RemovePanelGuaranteeViolation(violationInfo.PanelID); err != nil {
			return err
		}

		if panel.GuaranteeEndDate.Before(time.Now()) {
			panel.GuaranteeStatus = enum.PanelGuaranteeStatusExpired
		} else {
			panel.GuaranteeStatus = enum.PanelGuaranteeStatusActive
		}

		if err := installationService.installationRepository.UpdatePanel(tx, panel); err != nil {
			return err
		}
		return nil
	})

	return err
}

func (installationService *InstallationService) GetCorporationPanelGuaranteeViolation(violationInfo installationdto.GetCorporationGuaranteeViolationRequest) (guaranteedto.CorporationGuaranteeViolationResponse, error) {
	var violation guaranteedto.CorporationGuaranteeViolationResponse

	err := installationService.corporationService.CheckApplicantAccess(violationInfo.CorporationID, violationInfo.OperatorID)
	if err != nil {
		return violation, err
	}

	panel, err := installationService.getCorporationPanel(violationInfo.PanelID, violationInfo.CorporationID)
	if err != nil {
		return violation, err
	}

	if panel.GuaranteeStatus == enum.PanelGuaranteeStatusEmpty {
		return violation, exception.NotFoundError{Item: installationService.constants.Field.Guarantee}
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusVoided {
		return violation, exception.NotFoundError{Item: installationService.constants.Field.GuaranteeViolation}
	}

	violation, err = installationService.guaranteeService.GetCorporationPanelGuaranteeViolation(violationInfo.PanelID)
	if err != nil {
		return violation, err
	}

	return violation, nil
}

func (installationService *InstallationService) GetCustomerPanelGuaranteeViolation(violationInfo installationdto.GetCustomerGuaranteeViolationRequest) (guaranteedto.CustomerGuaranteeViolationResponse, error) {
	var violation guaranteedto.CustomerGuaranteeViolationResponse

	panel, err := installationService.getCustomerPanel(violationInfo.PanelID, violationInfo.OwnerID)
	if err != nil {
		return violation, err
	}

	if panel.GuaranteeStatus == enum.PanelGuaranteeStatusEmpty {
		return violation, exception.NotFoundError{Item: installationService.constants.Field.Guarantee}
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusVoided {
		return violation, exception.NotFoundError{Item: installationService.constants.Field.GuaranteeViolation}
	}

	violation, err = installationService.guaranteeService.GetCustomerPanelGuaranteeViolation(violationInfo.PanelID)
	if err != nil {
		return violation, err
	}

	return violation, nil
}

func (installationService *InstallationService) UpdatePanelGuaranteeViolation(violationInfo installationdto.UpdateGuaranteeViolationRequest) error {
	err := installationService.corporationService.CheckApplicantAccess(violationInfo.CorporationID, violationInfo.OperatorID)
	if err != nil {
		return err
	}

	panel, err := installationService.getCorporationPanel(violationInfo.PanelID, violationInfo.CorporationID)
	if err != nil {
		return err
	}

	if panel.GuaranteeStatus == enum.PanelGuaranteeStatusEmpty {
		return exception.NotFoundError{Item: installationService.constants.Field.Guarantee}
	}

	if panel.GuaranteeStatus != enum.PanelGuaranteeStatusVoided {
		return exception.NotFoundError{Item: installationService.constants.Field.GuaranteeViolation}
	}

	request := guaranteedto.UpdateGuaranteeViolationRequest{
		PanelID:    violationInfo.PanelID,
		OperatorID: violationInfo.OperatorID,
		Reason:     violationInfo.Reason,
		Details:    violationInfo.Details,
	}
	if err := installationService.guaranteeService.UpdateGuaranteeViolation(request); err != nil {
		return err
	}

	return nil
}
