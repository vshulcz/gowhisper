package services

import (
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"

	"github.com/google/uuid"
)

type MessageService struct {
	messageRepo repositories.MessageRepository
}

func NewMessageService(messageRepo repositories.MessageRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (s *MessageService) CreateMessage(content string, chatID uuid.UUID, userID uuid.UUID) (*entities.Message, error) {
	newMessage, _ := entities.NewMessage(content, chatID, userID)
	return newMessage, s.messageRepo.Save(newMessage)
}
