package models

// AddressFields struct
type AddressFields struct {
	AddressLine1 string `json:"addressLine1" bson:"addressLine1" validate:"required"`
	AddressLine2 string `json:"addressLine2" bson:"addressLine2"`
	LandMark     string `json:"landmark" bson:"landmark"`
	Village      string `json:"village" bson:"village" validate:"required"`
	District     string `json:"district" bson:"district" validate:"required"`
	State        string `json:"state" bson:"state" validate:"required"`
	Country      string `json:"country" bson:"country" validate:"required"`
	Pincode      string `json:"pincode" bson:"pincode" validate:"required,len=6"`
	Latitude     string `json:"latitude" bson:"latitude"`
	Longitude    string `json:"longitude" bson:"longitude"`
}
