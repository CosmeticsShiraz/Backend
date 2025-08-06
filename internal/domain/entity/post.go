package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type Post struct {
	database.Model
	Title         string      `json:"title"`
	CoverImage    string      `gorm:"type:varchar(255);default:null"`
	Content       string      `json:"content_html"`
	Description   string      `json:"description"`
	AuthorID      uint        `gorm:"not null;index"`
	Author        User        `gorm:"foreignKey:AuthorID"`
	CorporationID uint        `gorm:"not null;index"`
	Corporation   Corporation `gorm:"foreignKey:CorporationID"`
	Media         []Media     `gorm:"polymorphic:Owner;polymorphicValue:posts"`
	Likes         []Like      `gorm:"polymorphic:Owner;polymorphicValue:posts"`
	Status        enum.PostStatus
}
