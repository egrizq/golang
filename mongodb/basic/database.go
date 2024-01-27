package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
1.	context.Background() returns an empty Context. A Context in Go is a way to carry deadlines,
	cancellations, and other request-scoped values across API boundaries and between processes.
2. 	bson stands for Binary JSON, and it is a binary representation of JSON-like documents.
	BSON is used as the primary data format in MongoDB to store and exchange data between the MongoDB server
	and client applications.
*/

func main() {
	// todo connect with mongodb
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Panic(err)
	}

	// todo select the collection of database
	collection := client.Database("firstDB").Collection("users")

	// todo create an document
	document := map[string]interface{}{
		"name": "rizq",
		"age":  24,
	}

	// query to insert one document
	result, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Inserted document:", result.InsertedID)

	// todo query to show all the document
	find := bson.D{{}}
	cursor, err := collection.Find(context.Background(), &find)
	if err != nil {
		log.Panic(err)
	}
	defer cursor.Close(context.Background())

	// iterate through cursor and print all the data
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Panic(err)
		}

		// result[dataName] if want to selected data
		log.Println("All document from users:", result)
	}

	// todo update document
	document = map[string]interface{}{"nama": "rizq"}
	updatedDocument := map[string]interface{}{"$set": map[string]interface{}{"name": "rizq", "age": 20}}

	_, err = collection.UpdateOne(context.Background(), document, updatedDocument)
	if err != nil {
		panic(err)
	}
	log.Println("Successful updated document")

	// todo delete document
	documentDelete := bson.D{{Key: "nama", Value: "rizq"}}
	_, err = collection.DeleteOne(context.Background(), documentDelete)
	if err != nil {
		panic(err)
	}
	log.Println("Successful deleted document")

	// todo disconnect from mongodb
	if err = client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Disconnected from mongodb")
}
