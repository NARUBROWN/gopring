package web

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type ReturnValueHandler interface {
	Supports(t reflect.Type) bool
	Handle(value reflect.Value, c echo.Context) error
}
