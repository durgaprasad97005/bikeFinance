package response

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"github.com/durgaprasad97005/bikeFinance/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID                 bson.ObjectID      `json:"_id,omitempty"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	Email              string             `json:"email"`
	Phone              string             `json:"phone"`
	Role               constants.UserRole `json:"role"`
	Branch             constants.Branch   `json:"branch"`
	models.AuditFields `json:",inline"`
}
