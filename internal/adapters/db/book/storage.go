package book

type bookStorage struct {
}

func (bs *bookStorage) GetOne(uuid string) *Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) *Book {
	return nil
}
func (bs *bookStorage) Create(book *Book) *Book {
	return nil
}
func (bs *bookStorage) Delete(book *Book) error {
	return nil
}
