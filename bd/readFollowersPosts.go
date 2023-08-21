package bd

import (
	"context"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadFollowersPosts reads the posts of the followers*/
func ReadFollowersPosts(ID string, page int) ([]models.ReturnFollowersPosts, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("gatorNews")
	col := db.Collection("relationship")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "news",
			"localField":   "userrelationshipid",
			"foreignField": "userid",
			"as":           "news",
		}})

	conditions = append(conditions, bson.M{"$unwind": "$news"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"News.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnFollowersPosts
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
