package web

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type EchoContextResolver struct{}

func (r *EchoContextResolver) Supports(t reflect.Type) bool {
	return t == reflect.TypeOf((*echo.Context)(nil)).Elem()
}

func (r *EchoContextResolver) Resolve(rc *RequestContext, t reflect.Type, paramIndex int) (reflect.Value, error) {
	return reflect.ValueOf(rc.Echo), nil
}
