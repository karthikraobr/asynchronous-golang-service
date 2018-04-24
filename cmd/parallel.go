package main

import (
	"net/http"

	utils "github.com/karthikraobr/asynchronous-golang-service/utils"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", welcomePage)

	e.POST("/upload", utils.HandlePostRequest)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the matrix!")
}

func handleError(err error, context echo.Context) error {
	if err != nil {
		return context.String(http.StatusNotFound, err.Error())
	}
	return nil
}
