package web

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type ErrorReturnValueHandler struct{}

func (h *ErrorReturnValueHandler) Supports(t reflect.Type) bool {
	return t == reflect.TypeOf((*error)(nil)).Elem()
}

func (h *ErrorReturnValueHandler) Handle(v reflect.Value, c echo.Context) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}
