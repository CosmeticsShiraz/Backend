package installationdto

import (
	"time"

	addressdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/address"
	guaranteedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/guarantee"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
)

type NewInstallationRequest struct {
	OwnerID      uint
	Name         string
	Area         uint
	Power        uint
	MaxCost      float64
	BuildingType uint
	Description  string
	Address      addressdto.CreateAddressRequest
}

type UpdateInstallationRequest struct {
	RequestID    uint
	Name         *string
	Area         *uint
	Power        *uint
	MaxCost      *float64
	BuildingType *uint
	Status       *uint
	Description  *string
}

type CustomerRequestsListRequest struct {
	OwnerID uint
	Status  uint
	Offset  int
	Limit   int
}

type AdminInstallationListRequest struct {
	Status uint
	Offset int
	Limit  int
}

type CorporationPanelListRequest struct {
	CorporationID uint
	OperatorID    uint
	Status        uint
	Offset        int
	Limit         int
}

type CorporationPanelRequest struct {
	CorporationID  uint
	OperatorID     uint
	InstallationID uint
}

type CustomerPanelListRequest struct {
	OwnerID uint
	Status  uint
	Offset  int
	Limit   int
}

type GetOwnerRequest struct {
	OwnerID        uint
	InstallationID uint
}

type ChangeRequestStatusRequest struct {
	OwnerID   uint
	Status    enum.InstallationRequestStatus
	RequestID uint
}

type CompleteBidInstallationRequest struct {
	CorporationID   uint
	OperatorID      uint
	PanelID         uint
	Tilt            uint
	Azimuth         uint
	NumberOfModules uint
}

type AddPanelRequest struct {
	CorporationID        uint
	OperatorID           uint
	Name                 string
	Status               enum.PanelStatus
	CustomerPhone        string
	Power                uint
	Area                 uint
	BuildingType         uint
	Tilt                 uint
	Azimuth              uint
	TotalNumberOfModules uint
	GuaranteeID          *uint
	GuaranteeStartDate   *time.Time
	Address              addressdto.CreateAddressRequest
}

type CreateViolatePanelGuaranteeRequest struct {
	CorporationID      uint
	OperatorID         uint
	PanelID            uint
	GuaranteeViolation guaranteedto.CreateGuaranteeViolationRequest
}

type GetCorporationGuaranteeViolationRequest struct {
	CorporationID uint
	OperatorID    uint
	PanelID       uint
}

type GetCustomerGuaranteeViolationRequest struct {
	OwnerID uint
	PanelID uint
}

type UpdateGuaranteeViolationRequest struct {
	CorporationID uint
	OperatorID    uint
	PanelID       uint
	Reason        *string
	Details       *string
}

type UpdatePanelRequest struct {
	PanelID              uint
	Name                 *string
	Status               *uint
	BuildingType         *uint
	Area                 *uint
	Power                *uint
	Tilt                 *uint
	Azimuth              *uint
	TotalNumberOfModules *uint
}
