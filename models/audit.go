package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AuditFields struct {
	CreatedBy      bson.ObjectID `json:"createdBy" bson:"createdBy"`
	LastModifiedBy bson.ObjectID `json:"lastModifiedBy" bson:"lastModifiedBy"`
	CreatedAt      time.Time     `json:"createdAt" bson:"createdAt"`
	LastModifiedAt time.Time     `json:"lastModifiedAt" bson:"lastModifiedAt"`
}
