package web

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type StringReturnValueHandler struct{}

func (h *StringReturnValueHandler) Supports(t reflect.Type) bool {
	return t.Kind() == reflect.String
}

func (h *StringReturnValueHandler) Handle(v reflect.Value, c echo.Context) error {
	return c.String(200, v.String())
}
