package models

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Dealer struct
type Dealer struct {
	ID            bson.ObjectID    `bson:"_id,omitempty"`
	DealerCode    string           `bson:"dealerCode"`
	BusinessName  string           `bson:"businessName"`
	OwnerName     string           `bson:"ownerName"`
	Phone         string           `bson:"phone"`
	Email         string           `bson:"email"`
	GSTNumber     string           `bson:"gstNumber"`
	PANNumber     string           `bson:"panNumber"`
	AadhaarNumber string           `bson:"aadhaarNumber"`
	BranchID      string           `bson:"branchId"`
	Status        constants.Status `bson:"status"` 
	IsActive      bool             `bson:"isActive"`
	Address       AddressFields
	BankDetails   BankFields
	AuditFields   `bson:",inline"`
}
