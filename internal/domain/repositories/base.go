package repositories

type BaseRepository[T any] interface {
	Save(entity *T) error
	FindByID(id string) (*T, error)
	FindAll() ([]T, error)
	DeleteByID(id string) error
}
