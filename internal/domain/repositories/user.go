package repositories

import entities "gowhisper/internal/domain/entitites"

type UserRepository interface {
	BaseRepository[entities.User]
	FindByUsername(username string) (*entities.User, error)
}
