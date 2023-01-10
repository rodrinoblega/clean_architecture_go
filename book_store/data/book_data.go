package data

import (
	"context"
	"fmt"
	"github.com/rodrinoblega/clean_architecture_go/book_store/intrfcs"
)

type BookDataImpl struct {
	//DBConn *sql.DB
}

func NewBookDataImpl() intrfcs.BookData {
	return &BookDataImpl{}
}

func (bookDataImpl *BookDataImpl) GetBookByID(ctx context.Context, bookID int) {

}

func (bookDataImpl *BookDataImpl) TestBookService(ctx context.Context) {
	fmt.Println("*** Inside Book Data ***")
}
