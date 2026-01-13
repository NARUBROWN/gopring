package web

import (
	"fmt"
	"reflect"
	"strconv"
)

type RequestParamResolver struct{}

func (r *RequestParamResolver) Supports(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Int:
		return true
	default:
		return false
	}
}

func (r *RequestParamResolver) Resolve(rc *RequestContext, t reflect.Type, paramIndex int) (reflect.Value, error) {
	query := rc.Echo.QueryParams()
	if len(query) <= paramIndex {
		return reflect.Value{}, fmt.Errorf("query param이 부족합니다.")
	}

	i := 0
	for _, values := range query {
		if i == paramIndex {
			val := values[0]

			switch t.Kind() {
			case reflect.String:
				return reflect.ValueOf(val), nil
			case reflect.Int:
				n, err := strconv.Atoi(val)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(n), nil
			}
		}
		i++
	}

	return reflect.Value{}, fmt.Errorf("Query Param이 존재하지 않습니다 : %v", t)
}
