package entities

import (
	"errors"
)

type Chat struct {
	BaseEntity
	Users []*User
}

func NewChat(users []*User) (*Chat, error) {
	chat := &Chat{
		BaseEntity: *NewBaseEntity(),
		Users:      users,
	}

	if err := chat.Validate(); err != nil {
		return nil, err
	}
	return chat, nil
}

func (c *Chat) Validate() error {
	if len(c.Users) < 2 {
		return errors.New("a chat must have at least two users")
	}
	return nil
}
