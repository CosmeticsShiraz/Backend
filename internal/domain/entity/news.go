package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type News struct {
	database.Model
	Title       string  `json:"title"`
	Content     string  `json:"content_html"`
	Description string  `json:"description"`
	AuthorID    uint    `gorm:"not null;index"`
	Author      User    `gorm:"foreignKey:AuthorID"`
	CoverImage  string  `gorm:"type:text;default:null"`
	Media       []Media `gorm:"polymorphic:Owner;polymorphicValue:news"`
	Likes       []Like  `gorm:"polymorphic:Owner;polymorphicValue:news"`
	Status      enum.NewsStatus
}
