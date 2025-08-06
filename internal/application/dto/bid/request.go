package biddto

import (
	"time"

	paymentdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/payment"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
)

type SetBidRequest struct {
	CorporationID    uint
	RequestID        uint
	BidderID         uint
	Status           enum.BidStatus
	Cost             uint
	Area             uint
	Power            uint
	Description      string
	InstallationTime time.Time
	GuaranteeID      *uint
	PaymentTerms     paymentdto.PaymentTermsRequest
}

type UpdateBidRequest struct {
	CorporationID    uint
	BidID            uint
	BidderID         uint
	Cost             *uint
	Area             *uint
	Power            *uint
	Description      *string
	InstallationTime *time.Time
	GuaranteeID      *uint
	PaymentTerms     *paymentdto.UpdatePaymentTermsRequest
}

type GetBidRequest struct {
	CorporationID uint
	UserID        uint
	BidID         uint
}

type GetCorporationBidsRequest struct {
	CorporationID uint
	UserID        uint
	Status        uint
	Offset        int
	Limit         int
}

type GetListRequestBidsRequest struct {
	RequestID uint
	UserID    uint
	Offset    int
	Limit     int
}

type GetListRequestBidsRequestByAdmin struct {
	RequestID uint
	Offset    int
	Limit     int
}

type GetCustomerBidRequest struct {
	RequestID uint
	BidID     uint
	UserID    uint
}

type BidNotificationData struct {
	RequestID uint
	BidID     uint
	UserID    uint
}
