package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnFollowersPosts struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID             string             `bson:"userid" json:"userID,omitempty`
	UserRelationshipID string             `bson:"userrelationshipid" json:"userRelationshipID,omitempty"`
	News               struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
