package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endponint struct {
	s Service
}

func New(s Service) *Endponint {
	return &Endponint{
		s: s,
	}
}

func (e *Endponint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days Left: %d", d)

	err := ctx.String(http.StatusOK, s)

	if err != nil {
		return err
	}

	return nil
}
