package request

type RegisterUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,len=10"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role"`
	Branch    string `json:"branch"`
}

type LoginUser struct {
	Email     string `json:"email" validate:"required,email"`
	// Phone     string `json:"phone" validate:"required,len=10"`
	Password  string `json:"password" validate:"required"`
}