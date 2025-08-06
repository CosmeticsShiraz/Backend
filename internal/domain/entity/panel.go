package entity

import (
	"time"

	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Panel struct {
	database.Model
	Name                 string                    `gorm:"type:varchar(50);not null"`
	Status               enum.PanelStatus          `gorm:"not null"`
	BuildingType         enum.BuildingType         `gorm:"not null"`
	Area                 uint                      `gorm:"not null"`
	Power                uint                      `gorm:"not null"`
	Tilt                 uint                      `gorm:"not null"`
	Azimuth              uint                      `gorm:"not null"`
	TotalNumberOfModules uint                      `gorm:"not null"`
	CorporationID        uint                      `gorm:"not null;index"`
	Corporation          Corporation               `gorm:"foreignKey:CorporationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OperatorID           uint                      `gorm:"not null;index"`
	Operator             User                      `gorm:"foreignKey:OperatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CustomerID           uint                      `gorm:"not null;index"`
	Customer             User                      `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Address              Address                   `gorm:"polymorphic:Owner;polymorphicValue:panels"`
	GuaranteeStatus      enum.PanelGuaranteeStatus `gorm:"not null"`
	GuaranteeID          *uint                     `gorm:"index"`
	Guarantee            *Guarantee                `gorm:"foreignKey:GuaranteeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	GuaranteeStartDate   *time.Time
	GuaranteeEndDate     *time.Time
}
