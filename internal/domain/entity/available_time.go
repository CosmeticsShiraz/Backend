package entity

import (
	"time"

	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type AvailableTime struct {
	database.Model
	BidID     uint      `gorm:"not null;index"`
	Bid       Bid       `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
}
