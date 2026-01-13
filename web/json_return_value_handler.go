package web

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type JSONReturnValueHandler struct{}

func (h *JSONReturnValueHandler) Supports(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func (h *JSONReturnValueHandler) Handle(v reflect.Value, c echo.Context) error {
	return c.JSON(200, v.Interface())
}
