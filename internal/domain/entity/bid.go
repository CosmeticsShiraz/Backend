package entity

import (
	"time"

	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Bid struct {
	database.Model
	CorporationID    uint                `gorm:"index"`
	Corporation      Corporation         `gorm:"foreignKey:CorporationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BidderID         uint                `gorm:"index"`
	Bidder           User                `gorm:"foreignKey:BidderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RequestID        uint                `gorm:"not null;index"`
	Request          InstallationRequest `gorm:"constraint:OnDelete:CASCADE;"`
	Status           enum.BidStatus      `gorm:"not null;index"`
	Cost             uint                `gorm:"not null"`
	Area             uint                `gorm:"not null"`
	Power            uint                `gorm:"not null"`
	Description      string              `gorm:"type:text"`
	InstallationTime time.Time           `gorm:"not null"`
	PaymentTermsID   uint                `gorm:"index"`
	PaymentTerms     PaymentTerm         `gorm:"foreignKey:PaymentTermsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GuaranteeID      *uint               `gorm:"index"`
	Guarantee        *Guarantee          `gorm:"foreignKey:GuaranteeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	// AvailableTimes   []AvailableTime     `gorm:"foreignKey:BidID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
