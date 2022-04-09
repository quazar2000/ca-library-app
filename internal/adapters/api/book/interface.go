package book

type Service interface {
	GetBookByUUID(ctx context.Context, uuid string)
	GetAllBooks(ctx context.Context, limit, offset int)
	CreateBook(ctx context.Context, dto *CreateBookDTO) *Book
}