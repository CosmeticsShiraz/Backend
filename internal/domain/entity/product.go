package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Product struct {
	database.Model
	Name        string     `gorm:"type:varchar(100);not null;index"`
	Description string     `gorm:"type:text"`
	Price       int        `gorm:"not null"`
	Inventory   int        `gorm:"not null"`
	CategoryID  uint       `gorm:"index"`
	Category    Category   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandID     uint       `gorm:"index"`
	Brand       Brand      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pictures    []Picture  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}
