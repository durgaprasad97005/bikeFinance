package models

// AddressFields struct
type AddressFields struct {
	AddressLine1 string `json:"addressLine1" bson:"addressLine1"`
	AddressLine2 string `json:"addressLine2" bson:"addressLine2"`
	LandMark     string `json:"landmark" bson:"landmark"`
	Village      string `json:"village" bson:"village"`
	District     string `json:"district" bson:"district"`
	State        string `json:"state" bson:"state"`
	Country      string `json:"country" bson:"country"`
	Pincode      string `json:"pincode" bson:"pincode"`
	Latitude     string `json:"latitude" bson:"latitude"`
	Longitude    string `json:"longitude" bson:"longitude"`
}
