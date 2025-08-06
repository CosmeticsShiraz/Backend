package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type CorporationStaff struct {
	database.Model
	StaffID       uint           `gorm:"not null;index"`
	CorporationID uint           `gorm:"not null;index"`
	StaffType     enum.StaffType `gorm:"not null;index"`
}
