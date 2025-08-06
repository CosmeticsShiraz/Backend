package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Report struct {
	database.Model
	Description    string            `gorm:"type:text;not null"`
	ObjectID       uint              `gorm:"not null;index"`
	ObjectType     string            `gorm:"type:varchar(50);not null"`
	ReportedByID   uint              `gorm:"not null;index"`
	ReportedByType string            `gorm:"not null"`
	Status         enum.ReportStatus `gorm:"not null"`
}
