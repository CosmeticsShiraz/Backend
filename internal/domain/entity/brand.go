package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Brand struct {
	database.Model
	Name        string    `gorm:"type:varchar(100);not null"`
	Slug        string    `gorm:"type:varchar(100);uniqueIndex"`
	LogoPath    string    `gorm:"type:varchar(255);default:null"`
	Description string    `gorm:"type:text"`
	Website     string    `gorm:"type:varchar(255)"`
	Country     string    `gorm:"type:varchar(100)"`
	Products    []Product
}
