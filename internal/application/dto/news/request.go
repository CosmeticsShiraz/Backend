package newsdto

import (
	"mime/multipart"

	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
)

type CreateNewsRequest struct {
	Title       string
	Content     string
	Description string
	AuthorID    uint
	Status      enum.NewsStatus
	CoverImage  *multipart.FileHeader
}

type EditNewsRequest struct {
	NewsID      uint
	AuthorID    uint
	Title       *string
	Content     *string
	Description *string
	Status      uint
	CoverImage  *multipart.FileHeader
}

type EditNewsStatusRequest struct {
	NewsID   uint
	AuthorID uint
	Status   uint
}

type DeleteNewsRequest struct {
	NewsIDs  []uint
	AuthorID uint
}

type GetAdminNewsListRequest struct {
	Status uint
	Offset int
	Limit  int
}

type GetPublicNewsListRequest struct {
	Offset int
	Limit  int
}

// type GetNewsRequest struct {
// 	NewsID   uint
// 	UserType enum.UserType
// }

type AddNewsMediaRequest struct {
	NewsID   uint
	AuthorID uint
	Media    *multipart.FileHeader
}

type AccessMediaRequest struct {
	NewsID   uint
	AuthorID uint
	MediaID  uint
	UserType enum.UserType
}
