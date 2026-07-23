package models

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Dealer model struct
type Dealer struct {
	// Have to add the validate struct tag
	ID            bson.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	DealerCode    string           `json:"dealerCode" bson:"dealerCode"` // Auto generated field
	BusinessName  string           `json:"businessName" bson:"businessName"`
	OwnerName     string           `json:"ownerName" bson:"ownerName"`
	Phone         string           `json:"phone" bson:"phone"`
	Email         string           `json:"email" bson:"email"`
	GSTNumber     string           `json:"gstNumber" bson:"gstNumber"`
	PANNumber     string           `json:"panNumber" bson:"panNumber"`
	AadhaarNumber string           `json:"aadhaarNumber" bson:"aadhaarNumber"`
	Branch        string           `json:"branch" bson:"branch"`
	Status        constants.Status `json:"status" bson:"status"`
	IsActive      bool             `json:"isActive" bson:"isActive"`
	Address       AddressFields    `json:"address" bson:"address"`
	BankDetails   BankFields       `json:"bankDetails" bson:"bankDetails"`
	AuditFields   `bson:",inline"`
}
