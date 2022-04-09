package author

type Storage interface {
	GetOne(uuid string) *Author
	GetAll(limit, offset int) *Author
	Create(Author *Author) *Author
	Delete(Author *Author) error
}
