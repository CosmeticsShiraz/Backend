package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Category struct {
	database.Model
	Name        string     `gorm:"type:varchar(100);not null"`
	Slug        string     `gorm:"type:varchar(100);uniqueIndex"`
	Description string     `gorm:"type:text"`
	ParentID    *uint      `gorm:"index"`
	Parent      *Category  `gorm:"foreignKey:ParentID;constraint:OnDelete:SET NULL;"`
	Children    []Category `gorm:"foreignKey:ParentID"`
	Products    []Product
}
