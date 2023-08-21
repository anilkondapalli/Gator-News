package bd

import (
	"context"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*This is the final stop with the db to insert the user data*/
func InsertRegistration(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	u.ID = primitive.NewObjectID()
	
	db := MongoCN.Database("gatorNews")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
