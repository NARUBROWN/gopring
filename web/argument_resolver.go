package web

import (
	"reflect"
)

type ArgumentResolver interface {
	Supports(param reflect.Type) bool
	Resolve(rc *RequestContext, t reflect.Type, paramIndex int) (reflect.Value, error)
}
