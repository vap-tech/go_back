package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()

	mes := fmt.Sprintf("Количество дней до даты X: %d", d)

	err := ctx.String(http.StatusOK, mes)
	if err != nil {
		return err
	}
	return nil
}
