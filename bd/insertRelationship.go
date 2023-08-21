package bd

import (
	"context"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
)

/* This func inserts the relationship */
func InsertRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gatorNews")
	col := db.Collection("relationship")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
