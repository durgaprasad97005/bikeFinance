package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	FirstName    string        `bson:"firstName"`
	LastName     string        `bson:"lastName"`
	Email        string        `bson:"email"`
	Phone        string        `bson:"phone"`
	PasswordHash string        `bson:"passwordHash"`
	Role         string        `bson:"role"`
	Branch       string        `bson:"branch"`
	AuditFields  `bson:",inline"`
}
