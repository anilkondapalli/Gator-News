package bd

import (
	"context"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultRelation checks the relationship between 2 users*/
func ConsultRelation(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gatorNews")
	col := db.Collection("relationship")

	condition := bson.M{
		"userid":             t.UserID,
		"userrelationshipid": t.UserRelationshipID,
	}

	var result models.Relationship
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
