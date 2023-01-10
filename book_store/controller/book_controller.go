package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rodrinoblega/clean_architecture_go/book_store/intrfcs"
)

type BookController struct {
	bookService intrfcs.BookService
}

func NewBookController(echoCtx *echo.Echo, bookService intrfcs.BookService) {
	bookController := &BookController{
		bookService: bookService,
	}

	echoCtx.GET("/printAuthor", bookController.PrintAuthor)
	echoCtx.GET("test", bookController.Test)
}

func (bookController *BookController) PrintAuthor(echoCtx echo.Context) error {
	return nil
}

func (bookController *BookController) Test(ec echo.Context) error {
	fmt.Println("*** Inside Book Controller ***")

	requestContext := ec.Request().Context()
	bookController.bookService.TestBookService(requestContext)
	return nil
}
