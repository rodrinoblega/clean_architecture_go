package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rodrinoblega/clean_architecture_go/book_store/controller"
	"github.com/rodrinoblega/clean_architecture_go/book_store/data"
	"github.com/rodrinoblega/clean_architecture_go/book_store/service"
)

func main() {
	echoContext := echo.New()

	data := data.NewBookDataImpl()
	service := service.NewBookServiceImpl(data)

	controller.NewBookController(echoContext, service)

	echoContext.Logger.Info(echoContext.Start(":8090"))
}
