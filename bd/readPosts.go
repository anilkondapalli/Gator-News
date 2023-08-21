package bd

import (
	"context"
	"log"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ReadPosts reads the posts of a profile*/
func ReadPosts(ID string, page int64) ([]*models.ReturnPost, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("gatorNews")
	col := db.Collection("news")

	var result []*models.ReturnPost

	if ID == "" {
		log.Println("userid provided is empty!")
		return result, false
	}
	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var record models.ReturnPost
		err := cursor.Decode(&record)
		if err != nil {
			return result, false
		}
		result = append(result, &record)

	}
	return result, true

}
