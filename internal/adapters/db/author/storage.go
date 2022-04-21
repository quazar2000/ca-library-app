package author

import "github.com/quazar2000/ca-library-app/internal/domain/author"

type authorStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) author.Storage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (bs *authorStorage) GetAll(limit, offset int) *author.Author {
	return nil
}
func (bs *authorStorage) Create(author *author.Author) *author.Author {
	return nil
}
func (bs *authorStorage) Delete(author *author.Author) error {
	return nil
}
