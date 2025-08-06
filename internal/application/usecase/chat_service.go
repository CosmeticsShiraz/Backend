package usecase

import (
	chatdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/chat"
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
)

type ChatService interface {
	CreateChatRoom(request chatdto.CreateOrGetUserRoomRequest) (*entity.ChatRoom, error)
	CreateOrGetRoom(request chatdto.CreateOrGetUserRoomRequest) (chatdto.ChatRoomDetailsResponse, error)
	GetCorporationRoom(request chatdto.GetCorporationRoomRequest) (chatdto.ChatRoomDetailsResponse, error)
	GetUserRooms(userID uint) ([]chatdto.ChatRoomDetailsResponse, error)
	GetCorporationRooms(request chatdto.GetCorporationRoomsRequest) ([]chatdto.ChatRoomDetailsResponse, error)
	SaveMessage(roomID, senderID uint, content string) (chatdto.RoomMessagesResponse, error)
	GetRoomMessages(request chatdto.GetRoomMessageRequest) ([]chatdto.RoomMessagesResponse, error)
	BlockChatRoom(request chatdto.BlockServiceChatRequest) error
	UnBlockChatRoom(request chatdto.BlockServiceChatRequest) error
}
