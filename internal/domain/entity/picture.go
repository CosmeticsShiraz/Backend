package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Picture struct {
	database.Model
	Path      string `gorm:"type:varchar(255);not null"`
	ProductID uint   `gorm:"index"`
}
