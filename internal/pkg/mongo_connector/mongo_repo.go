package mongo_connector

import (
	"context"
	"go-ddd-quickstart/internal/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	Collection *mongo.Collection
}

func (repo *MongoRepo) Create(item db.IItem) (string, error) {
	result, err := repo.Collection.InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil
}
func (repo *MongoRepo) Update(id string, item db.IItem) error {
	_, err := repo.Collection.UpdateByID(context.TODO(), id, item)
	if err != nil {
		return err
	}
	return nil
}
func (repo *MongoRepo) Delete(id string) error {
	_, err := repo.Collection.DeleteOne(context.TODO(), id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *MongoRepo) List(filter map[string]interface{}) ([]db.IItem, error) {
	cursor, err := repo.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var items []db.IItem
	for cursor.Next(context.TODO()) {
		var item db.IItem
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
func (repo *MongoRepo) Retrieve(id string) (db.IItem, error) {
	var item db.IItem
	err := repo.Collection.FindOne(context.TODO(), id).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
