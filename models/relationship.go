package models

/* Relationship model to record the relationship of one user with another */
type Relationship struct {
	UserID             string `bson:"userid" json:"userID"`
	UserRelationshipID string `bson:"userrelationshipid" json:"userRelationshipID"`
}
