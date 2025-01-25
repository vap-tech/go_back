package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
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

	d := time.Date(2025, time.April, 10, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	mes := fmt.Sprintf("Количество дней до даты X: %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, mes)
	if err != nil {
		return err
	}
	return nil
}
