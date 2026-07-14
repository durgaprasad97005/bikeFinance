package response

import (
	"github.com/durgaprasad97005/bikeFinance/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID                 bson.ObjectID `json:"_id,omitempty"`
	FirstName          string        `json:"firstName"`
	LastName           string        `json:"lastName"`
	Email              string        `json:"email"`
	Phone              string        `json:"phone"`
	Role               string        `json:"role"`
	Branch             string        `json:"branch"`
	models.AuditFields `json:",inline"`
}
