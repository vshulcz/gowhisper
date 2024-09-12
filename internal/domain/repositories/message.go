package repositories

import entities "gowhisper/internal/domain/entitites"

type MessageRepository interface {
	BaseRepository[entities.Message]
	FindByChatID(chatID string) ([]*entities.Message, error)
}
