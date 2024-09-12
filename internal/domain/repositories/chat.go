package repositories

import entities "gowhisper/internal/domain/entitites"

type ChatRepository interface {
	BaseRepository[entities.Chat]
}
