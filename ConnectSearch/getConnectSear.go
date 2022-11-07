package ConnectSearch

import "go.mongodb.org/mongo-driver/mongo"

func GetCollectionsearch(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Humans").Collection("Humans")
	return collection
}