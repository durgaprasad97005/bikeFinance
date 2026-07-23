package request

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"github.com/durgaprasad97005/bikeFinance/models"
)

// Dealer request struct
type DealerRequestDTO struct {
	BusinessName  string               `json:"businessName" validate:"required"`
	OwnerName     string               `json:"ownerName" validate:"required"`
	Phone         string               `json:"phone"`
	Email         string               `json:"email" validate:"required,email"`
	GSTNumber     string               `json:"gstNumber"`
	PANNumber     string               `json:"panNumber"`
	AadhaarNumber string               `json:"aadhaarNumber"`
	BranchID      string               `json:"branchId"`
	Status        constants.Status     `json:"status"`
	IsActive      bool                 `json:"isActive"`
	Address       models.AddressFields `json:"address"`
	BankDetails   models.BankFields    `json:"bankDetails"`
}
