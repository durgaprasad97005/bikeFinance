package models

import (
	"github.com/durgaprasad97005/bikeFinance/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID           bson.ObjectID      `bson:"_id,omitempty"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Email        string             `bson:"email"`
	Phone        string             `bson:"phone"`
	PasswordHash string             `bson:"passwordHash"`
	Role         constants.UserRole `bson:"role"`
	Branch       constants.Branch   `bson:"branch"`
	AuditFields  `bson:",inline"`
}
