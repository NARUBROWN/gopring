package web

import "reflect"

type HandlerMethod struct {
	Bean    any            // Controller 인스턴스
	Method  reflect.Method // 호출할 메서드
	Mapping RequestMapping
}
