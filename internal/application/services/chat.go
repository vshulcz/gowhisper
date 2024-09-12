package services

import (
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"
	"log"
)

type ChatService struct {
	chatRepo repositories.ChatRepository
}

func NewChatService(chatRepo repositories.ChatRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (s *ChatService) CreateChat(users []*entities.User) (*entities.Chat, error) {
	newChat, err := entities.NewChat(users)
	if err != nil {
		log.Printf("Error creating chat: %v", err)
		return nil, err
	}
	if err := s.chatRepo.Save(newChat); err != nil {
		log.Printf("Error saving chat: %v", err)
		return nil, err
	}
	return newChat, nil
}
