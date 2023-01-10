package intrfcs

import (
	"context"
	"github.com/rodrinoblega/clean_architecture_go/book_store/model"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
	TestBookService(ctx context.Context)
}

type BookData interface {
	GetBookByID(ctx context.Context, bookID int)
	TestBookService(ctx context.Context)
}
