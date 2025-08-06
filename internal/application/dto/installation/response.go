package installationdto

import (
	"time"

	addressdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/address"
	corporationdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/corporation"
	guaranteedto "github.com/CosmeticsShiraz/Backend/internal/application/dto/guarantee"
	userdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/user"
)

type EnumStatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AnonymousRequestsResponse struct {
	ID           uint                       `json:"id"`
	Name         string                     `json:"name"`
	CreatedTime  time.Time                  `json:"createdTime"`
	Status       string                     `json:"status"`
	PowerRequest uint                       `json:"powerRequest"`
	MaxCost      float64                    `json:"maxCost"`
	BuildingType string                     `json:"buildingType"`
	Address      addressdto.AddressResponse `json:"address"`
}

type PublicRequestDetailsResponse struct {
	ID           uint                       `json:"id"`
	Name         string                     `json:"name"`
	Status       string                     `json:"status"`
	PowerRequest uint                       `json:"powerRequest"`
	Description  string                     `json:"description"`
	BuildingType string                     `json:"buildingType"`
	Area         uint                       `json:"area"`
	MaxCost      float64                    `json:"maxCost"`
	Customer     userdto.CredentialResponse `json:"customer"`
	Address      addressdto.AddressResponse `json:"address"`
}

type CorporationPanelListResponse struct {
	ID                   uint                       `json:"id"`
	Name                 string                     `json:"name"`
	Status               string                     `json:"status"`
	BuildingType         string                     `json:"buildingType"`
	Area                 uint                       `json:"area"`
	Power                uint                       `json:"power"`
	Tilt                 uint                       `json:"tilt"`
	Azimuth              uint                       `json:"azimuth"`
	TotalNumberOfModules uint                       `json:"totalNumberOfModules"`
	GuaranteeStatus      string                     `json:"guaranteeStatus"`
	Operator             userdto.CredentialResponse `json:"operator"`
	Customer             userdto.CredentialResponse `json:"customer"`
	Address              addressdto.AddressResponse `json:"address"`
}

type CorporationPanelResponse struct {
	ID                   uint                           `json:"id"`
	Name                 string                         `json:"name"`
	Status               string                         `json:"status"`
	BuildingType         string                         `json:"buildingType"`
	Area                 uint                           `json:"area"`
	Power                uint                           `json:"power"`
	Tilt                 uint                           `json:"tilt"`
	Azimuth              uint                           `json:"azimuth"`
	TotalNumberOfModules uint                           `json:"totalNumberOfModules"`
	GuaranteeStatus      string                         `json:"guaranteeStatus"`
	Operator             userdto.CredentialResponse     `json:"operator"`
	Customer             userdto.CredentialResponse     `json:"customer"`
	Address              addressdto.AddressResponse     `json:"address"`
	Guarantee            guaranteedto.GuaranteeResponse `json:"guarantee"`
}

type AdminPanelResponse struct {
	ID                   uint                                         `json:"id"`
	Name                 string                                       `json:"name"`
	Status               string                                       `json:"status"`
	BuildingType         string                                       `json:"buildingType"`
	Area                 uint                                         `json:"area"`
	Power                uint                                         `json:"power"`
	Tilt                 uint                                         `json:"tilt"`
	Azimuth              uint                                         `json:"azimuth"`
	TotalNumberOfModules uint                                         `json:"totalNumberOfModules"`
	GuaranteeStatus      string                                       `json:"guaranteeStatus"`
	Operator             userdto.CredentialResponse                   `json:"operator"`
	Customer             userdto.CredentialResponse                   `json:"customer"`
	Corporation          corporationdto.CorporationCredentialResponse `json:"corporation"`
	Address              addressdto.AddressResponse                   `json:"address"`
	Guarantee            guaranteedto.GuaranteeResponse               `json:"guarantee"`
}

type CustomerPanelListResponse struct {
	ID                   uint                                         `json:"id"`
	Name                 string                                       `json:"name"`
	Status               string                                       `json:"status"`
	BuildingType         string                                       `json:"buildingType"`
	Area                 uint                                         `json:"area"`
	Power                uint                                         `json:"power"`
	Tilt                 uint                                         `json:"tilt"`
	Azimuth              uint                                         `json:"azimuth"`
	TotalNumberOfModules uint                                         `json:"totalNumberOfModules"`
	GuaranteeStatus      string                                       `json:"guaranteeStatus"`
	Corporation          corporationdto.CorporationCredentialResponse `json:"corporation"`
	Address              addressdto.AddressResponse                   `json:"address"`
}

type CustomerPanelResponse struct {
	ID                   uint                                         `json:"id"`
	Name                 string                                       `json:"name"`
	Status               string                                       `json:"status"`
	BuildingType         string                                       `json:"buildingType"`
	Area                 uint                                         `json:"area"`
	Power                uint                                         `json:"power"`
	Tilt                 uint                                         `json:"tilt"`
	Azimuth              uint                                         `json:"azimuth"`
	TotalNumberOfModules uint                                         `json:"totalNumberOfModules"`
	GuaranteeStatus      string                                       `json:"guaranteeStatus"`
	Corporation          corporationdto.CorporationCredentialResponse `json:"corporation"`
	Address              addressdto.AddressResponse                   `json:"address"`
	Guarantee            guaranteedto.GuaranteeResponse               `json:"guarantee"`
}
