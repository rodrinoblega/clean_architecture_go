package service

import (
	"context"
	"fmt"
	"github.com/rodrinoblega/clean_architecture_go/book_store/intrfcs"
	"github.com/rodrinoblega/clean_architecture_go/book_store/model"
)

type BookServiceImpl struct {
	bookData intrfcs.BookData
}

func NewBookServiceImpl(bookData intrfcs.BookData) intrfcs.BookService {
	return &BookServiceImpl{
		bookData: bookData,
	}
}

func (bookService *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book) {
	fmt.Println("asd")
}

func (bookServiceImpl *BookServiceImpl) TestBookService(ctx context.Context) {
	fmt.Println("*** Inside Book Service ***")
	bookServiceImpl.bookData.TestBookService(ctx)
}
