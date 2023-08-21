package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This function is for deleting posts
func DeletePost(ID string,UserID string) error{
	ctx,cancel:= context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()

	db := MongoCN.Database("gatorNews")
    col := db.Collection("news")

	objID,_ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":objID,
		"userid":UserID,
	}
	_,err := col.DeleteOne(ctx,condition)
	return err
}