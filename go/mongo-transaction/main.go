package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID     string
	Name   string
	Amount int
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to verify that the client can connect
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	// Create a new session for the transaction
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	// Define the transaction options
	opts := options.Transaction().SetMaxCommitTime(1000).SetUUID(uuid.New())

	// Define the transaction function
	txFunc := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		collection := client.Database("test").Collection("accounts")

		// Read account 1
		account1 := Account{}
		err := collection.FindOne(sessionCtx, bson.M{"id": "1"}).Decode(&account1)
		if err != nil {
			return nil, err
		}

		// Read account 2
		account2 := Account{}
		err = collection.FindOne(sessionCtx, bson.M{"id": "2"}).Decode(&account2)
		if err != nil {
			return nil, err
		}

		// Transfer funds from account 1 to account 2
		account1.Amount -= 100
		account2.Amount += 100

		// Update account 1
		_, err = collection.UpdateOne(sessionCtx, bson.M{"id": "1"}, bson.M{"$set": bson.M{"amount": account1.Amount}})
		if err != nil {
			return nil, err
		}

		// Update account 2
		_, err = collection.UpdateOne(sessionCtx, bson.M{"id": "2"}, bson.M{"$set": bson.M{"amount": account2.Amount}})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Start the transaction
	result, err := session.WithTransaction(context.Background(), txFunc, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction completed: %v\n", result)
}
