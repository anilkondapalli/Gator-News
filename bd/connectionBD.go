package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN is an object to conect to DB */
var MongoCN = ConnectorBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gatorUsers:Gators2023@gator-news.djwv2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* connectorBD allows the connection to DB*/
func ConnectorBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful Connection with DB")
	return client
}

/* Checks connection ping to DB */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
