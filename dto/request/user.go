package request

import "github.com/durgaprasad97005/bikeFinance/constants"

type RegisterUser struct {
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"required,email"`
	Phone     string             `json:"phone" validate:"required,len=10"`
	Password  string             `json:"password" validate:"required,min=8"`
	Role      constants.UserRole `json:"role"`
	Branch    constants.Branch   `json:"branch"`
}

type LoginUser struct {
	Email string `json:"email" validate:"required,email"`
	// Phone     string `json:"phone" validate:"required,len=10"`
	Password string `json:"password" validate:"required,min=8"`
}
