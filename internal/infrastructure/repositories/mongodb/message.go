package mongodb

import (
	"context"
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMessageRepository struct {
	*MongoRepository[entities.Message]
}

func NewMongoMessageRepository(db *mongo.Database) repositories.MessageRepository {
	return &MongoMessageRepository{
		MongoRepository: NewMongoRepository[entities.Message](db, "messages"),
	}
}

func (r *MongoMessageRepository) FindByChatID(chatID string) ([]*entities.Message, error) {
	var messages []*entities.Message
	cursor, err := r.collection.Find(context.Background(), bson.M{"chat_id": chatID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var message entities.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
