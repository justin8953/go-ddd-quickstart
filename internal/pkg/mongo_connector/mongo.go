package mongo_connector

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func NewMongoService(collections []string) *MongoService {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Error("Set your 'MONGODB_URI' environment variable.")
		panic("Set your 'MONGODB_URI' environment variable.")
	}
	db := os.Getenv("MONGODB_DB")
	if db == "" {
		log.Error("Set your 'MONGODB_DB' environment variable.")
		panic("Set your 'MONGODB_DB' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		log.Error("Failed to connect to MongoDB:", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Error("Failed to disconnect from MongoDB:", err)
		}
	}()

	collectionsMap := make(map[string]*mongo.Collection)
	for _, collection := range collections {
		collectionsMap[collection] = client.Database(db).Collection(collection)
	}
	return &MongoService{
		Client:      client,
		Collections: collectionsMap,
	}
}

func (s *MongoService) GetMongoRepo(collection string) *MongoRepo {
	return &MongoRepo{
		Collection: s.Collections[collection],
	}
}
