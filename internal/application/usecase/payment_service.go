package usecase

import paymentdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/payment"

type PaymentService interface {
	GetPaymentMethods() []paymentdto.PaymentMethodResponse
	GetPaymentTerms(payTermID uint) (paymentdto.PaymentTermsResponse, error)
	CreatePaymentTerms(paymentTermsRequest paymentdto.PaymentTermsRequest) (uint, error)
	UpdatePaymentTerms(updatePaymentRequest paymentdto.UpdatePaymentTermsRequest) error
}
