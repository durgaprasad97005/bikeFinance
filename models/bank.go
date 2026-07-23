package models

// Bank struct
type BankFields struct {
	BankName      string `json:"bankName" bson:"bankName" validate:"required"`
	BranchName    string `json:"branchName" bson:"branchName" validate:"required"`
	AccountHolder string `json:"accountHolder" bson:"accountHolder" validate:"required"`
	AccountNumber string `json:"accountNumber" bson:"accountNumber" validate:"required"`
	IFSCCode      string `json:"ifscCode" bson:"ifscCode" validate:"required"`
}
