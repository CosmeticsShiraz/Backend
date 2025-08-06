package entity

import (
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type ChatMessage struct {
	database.Model
	RoomID   uint   `json:"room_id"`
	SenderID uint   `json:"sender_id"`
	Content  string `json:"content"`
}
