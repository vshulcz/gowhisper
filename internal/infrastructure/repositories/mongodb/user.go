package mongodb

import (
	"context"
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	*MongoRepository[entities.User]
}

func NewMongoUserRepository(db *mongo.Database) repositories.UserRepository {
	return &MongoUserRepository{
		MongoRepository: NewMongoRepository[entities.User](db, "users"),
	}
}

func (r *MongoUserRepository) FindByUsername(username string) (*entities.User, error) {
	var result entities.User
	err := r.MongoRepository.GetCollection().FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
