package websocket

import (
	"encoding/json"
	"time"

	userdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/user"
)

// convert to enum
const (
	MessageTypeChat         = "chat"
	MessageTypeNotification = "notification"
)

type Message struct {
	MessageID uint                       `json:"id"`
	Sender    userdto.CredentialResponse `json:"sender"`
	Type      string                     `json:"type"`
	RoomID    uint                       `json:"room_id,omitempty"`
	SenderID  uint                       `json:"sender_id,omitempty"`
	Content   json.RawMessage            `json:"content"`
	Timestamp time.Time                  `json:"timestamp"`
	Client    *Client                    `json:"-"`
}

type NotificationPayload struct {
	ID             uint        `json:"id"`
	Type           string      `json:"type"`
	Description    string      `json:"description"`
	AdditionalData interface{} `json:"additionalData"`
	IsRead         bool        `json:"is_read"`
	CreatedAt      string      `json:"created_at"`
}
