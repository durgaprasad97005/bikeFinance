package request

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"github.com/durgaprasad97005/bikeFinance/models"
)

// Dealer request struct
type DealerRequestDTO struct {
	BusinessName  string               `json:"businessName" validate:"required"`
	OwnerName     string               `json:"ownerName" validate:"required"`
	Phone         string               `json:"phone" validate:"omitempty,len=10"`
	Email         string               `json:"email" validate:"required,email"`
	GSTNumber     string               `json:"gstNumber" validate:"omitempty,len=10"`
	PANNumber     string               `json:"panNumber" validate:"omitempty,len=10"`
	AadhaarNumber string               `json:"aadhaarNumber" validate:"omitempty,len=12,digits"`
	Branch        constants.Branch     `json:"branch"`
	Status        constants.Status     `json:"status"`
	IsActive      bool                 `json:"isActive"`
	Address       models.AddressFields `json:"address"`
	BankDetails   models.BankFields    `json:"bankDetails"`
}
