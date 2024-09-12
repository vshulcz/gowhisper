package entities

import (
	"errors"

	"github.com/google/uuid"
)

type Message struct {
	BaseEntity
	Content string
	ChatID  uuid.UUID
	UserID  uuid.UUID
}

func NewMessage(content string, chatID, userID uuid.UUID) (*Message, error) {
	message := &Message{
		BaseEntity: *NewBaseEntity(),
		Content:    content,
		ChatID:     chatID,
		UserID:     userID,
	}

	if err := message.Validate(); err != nil {
		return nil, err
	}
	return message, nil
}

func (m *Message) Validate() error {
	if m.Content == "" {
		return errors.New("message content cannot be empty")
	}
	return nil
}
