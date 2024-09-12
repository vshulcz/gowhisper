package mongodb

import (
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoChatRepository struct {
	*MongoRepository[entities.Chat]
}

func NewMongoChatRepository(db *mongo.Database) repositories.ChatRepository {
	return &MongoChatRepository{
		MongoRepository: NewMongoRepository[entities.Chat](db, "chats"),
	}
}
