package bd

import (
	"context"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*insertPost saves the post to the db*/

func InsertPost(t models.RecordPost) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("gatorNews")
	col := db.Collection("news")
	registration := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, registration)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
