package web

import (
	"gopring/context"
	"reflect"

	"github.com/labstack/echo/v4"
)

type ControllerMapping struct {
	Path    string
	Handler echo.HandlerFunc
}

func CreateControllerMappings(ctx *context.ApplicationContext) []HandlerMethod {
	var handlers []HandlerMethod

	for _, definition := range ctx.BeanDefinitions {
		beanType := definition.BeanType
		bean := ctx.GetBean(beanType)

		// struct 타입만 Controller 후보
		if beanType.Kind() != reflect.Ptr || beanType.Elem().Kind() != reflect.Struct {
			continue
		}

		// Mappings() 메서드가 없으면 생략
		mappingsMethod, ok := beanType.MethodByName("Mappings")
		if !ok {
			continue
		}

		// mappings := controller.Mappings()
		results := mappingsMethod.Func.Call(
			[]reflect.Value{reflect.ValueOf(bean)},
		)

		mappings := results[0].Interface().(map[string]RequestMapping)

		for methodName, mapping := range mappings {
			method, ok := beanType.MethodByName(methodName)
			if !ok {
				panic("해당하는 핸들러 메서드가 없습니다: " + methodName)
			}

			handlers = append(handlers, HandlerMethod{
				Bean:    bean,
				Method:  method,
				Mapping: mapping,
			})
		}
	}

	return handlers
}

func isValidHandlerMethod(m reflect.Method) bool {
	// 메서드 시그니처
	t := m.Type

	// receiver + echo.Context
	if t.NumIn() != 2 {
		return false
	}

	if t.In(1) != reflect.TypeOf((*echo.Context)(nil)).Elem() {
		return false
	}

	// 반환 에러
	if t.NumOut() != 1 {
		return false
	}

	if t.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return false
	}

	return true
}
