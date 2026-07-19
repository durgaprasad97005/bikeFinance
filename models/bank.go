package models

// Bank struct
type BankFields struct {
	BankName      string `json:"bankName" bson:"bankName"`
	BranchName    string `json:"branchName" bson:"branchName"`
	AccountHolder string `json:"accountHolder" bson:"accountHolder"`
	AccountNumber string `json:"accountNumber" bson:"accountNumber"`
	IFSCCode      string `json:"ifscCode" bson:"ifscCode"`
}
