package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running")

	// Echo instance
	s := echo.New()
	s.GET("/status", Handler)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Hello World")
	if err != nil {
		return err
	}
	return nil
}
