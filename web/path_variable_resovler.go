package web

import (
	"fmt"
	"reflect"
	"strconv"
)

type PathVariableResolver struct{}

func (r *PathVariableResolver) Supports(t reflect.Type) bool {
	return t.Kind() == reflect.Int || t.Kind() == reflect.String
}

func (r *PathVariableResolver) Resolve(rc *RequestContext, t reflect.Type, paramIndex int) (reflect.Value, error) {

	if len(rc.PathVars) == 0 {
		return reflect.Value{}, fmt.Errorf("path variable이 없습니다.")
	}

	// PathVars는 순서대로 정렬되어있다고 가정
	values := make([]string, 0, len(rc.PathVars))
	for _, v := range rc.PathVars {
		values = append(values, v)
	}

	if paramIndex >= len(values) {
		return reflect.Value{}, fmt.Errorf("path variable 부족")
	}

	raw := values[paramIndex]

	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(raw), nil
	case reflect.Int:
		n, err := strconv.Atoi(raw)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(n), nil
	}

	return reflect.Value{}, fmt.Errorf("지원하지 않는 타입")
}
