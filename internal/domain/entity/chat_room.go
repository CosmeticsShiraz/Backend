package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type ChatRoom struct {
	database.Model
	CorporationID uint            `gorm:"not null;index"`
	Corporation   Corporation     `gorm:"foreignKey:CorporationID"`
	CustomerID    uint            `gorm:"not null;index"`
	Customer      User            `gorm:"foreignKey:CustomerID"`
	Status        enum.ChatStatus `gorm:"default:1"`
	BlockedBy     *enum.BlockedBy
}
