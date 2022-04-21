package book

import "github.com/quazar2000/ca-library-app/internal/domain/book"

type bookStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) book.Storage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) *book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book.Book) error {
	return nil
}
