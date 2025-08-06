package corporationdto

import (
	addressdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/address"
	userdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/user"
)

type CorporationCredentialResponse struct {
	ID          uint                         `json:"id"`
	Name        string                       `json:"name"`
	Logo        string                       `json:"logo"`
	ContactInfo []ContactInformationResponse `json:"contactInfo"`
	Addresses   []addressdto.AddressResponse `json:"addresses"`
}

type CorporationPrivateInfoResponse struct {
	ID                     uint                         `json:"id"`
	Name                   string                       `json:"name"`
	RegistrationNumber     string                       `json:"registrationNumber"`
	NationalID             string                       `json:"nationalID"`
	IBAN                   string                       `json:"iban"`
	Logo                   string                       `json:"logo"`
	VATTaxpayerCertificate string                       `json:"vatTaxpayerCertificate"`
	OfficialNewspaperAD    string                       `json:"officialNewspaperAD"`
	Signatories            []SignatoryResponse          `json:"signatories"`
	ContactInfo            []ContactInformationResponse `json:"contactInfo"`
	Addresses              []addressdto.AddressResponse `json:"addresses"`
}

type SignatoryResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	NationalCardNumber string `json:"nationalCardNumber"`
	Position           string `json:"position"`
}

type ContactInformationResponse struct {
	ID          uint                `json:"id"`
	ContactType ContactTypeResponse `json:"contactType"`
	Value       string              `json:"value"`
}

type ContactTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetStatusesResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type GetAdminCorporationReview struct {
	Reviewer userdto.CredentialResponse `json:"reviewer"`
	Action   string                     `json:"action"`
	Reason   *string                    `json:"reason"`
	Notes    *string                    `json:"notes"`
}

type GetCustomerCorporationReview struct {
	Action string  `json:"action"`
	Reason *string `json:"reason"`
	Notes  *string `json:"notes"`
}
