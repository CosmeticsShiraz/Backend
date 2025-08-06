package postgres

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type InstallationRepository interface {
	FindRequestsByStatus(db database.Database, status []enum.InstallationRequestStatus, opts ...QueryModifier) ([]*entity.InstallationRequest, error)
	FindRequestByID(db database.Database, requestID uint) (*entity.InstallationRequest, error)
	FindRequestByOwner(db database.Database, requestID, ownerID uint) (*entity.InstallationRequest, error)
	FindOwnerRequests(db database.Database, ownerID uint, status []enum.InstallationRequestStatus, opts ...QueryModifier) ([]*entity.InstallationRequest, error)
	FindOwnerRequestByName(db database.Database, ownerID uint, status []enum.InstallationRequestStatus, name string) (*entity.InstallationRequest, error)
	FindCorporationPanel(db database.Database, panelID, corporationID uint) (*entity.Panel, error)
	FindCorporationPanels(db database.Database, corporationID uint, allowedStatus []enum.PanelStatus, opts ...QueryModifier) ([]*entity.Panel, error)
	FindCustomerPanels(db database.Database, customerID uint, allowedStatus []enum.PanelStatus, opts ...QueryModifier) ([]*entity.Panel, error)
	FindCustomerPanel(db database.Database, panelID, customerID uint) (*entity.Panel, error)
	CreateRequest(db database.Database, request *entity.InstallationRequest) error
	UpdateRequest(db database.Database, request *entity.InstallationRequest) error
	DeleteRequest(db database.Database, request *entity.InstallationRequest) error
	FindPanelByNameAndCustomerID(db database.Database, panelName string, customerID uint) (*entity.Panel, error)
	FindPanelByOwner(db database.Database, panelID, customerID uint) (*entity.Panel, error)
	FindPanelByID(db database.Database, panelID uint) (*entity.Panel, error)
	FindPanelsByStatus(db database.Database, allowedStatus []enum.PanelStatus, opts ...QueryModifier) ([]*entity.Panel, error)
	CreatePanel(db database.Database, panel *entity.Panel) error
	UpdatePanel(db database.Database, panel *entity.Panel) error
	DeletePanel(db database.Database, panel *entity.Panel) error
}
