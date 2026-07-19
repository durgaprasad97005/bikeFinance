package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Counter struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	DealerCode uint32        `bson:"dealerCode"`
}
