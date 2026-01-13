package web

import (
	"fmt"
	"gopring/context"
	"reflect"

	"github.com/labstack/echo/v4"
)

type RequestContext struct {
	Echo     echo.Context
	PathVars map[string]string
}

type Dispatcher struct {
	ctx            *context.ApplicationContext
	handlers       []HandlerMethod
	resolvers      []ArgumentResolver
	returnHandlers []ReturnValueHandler
}

func NewDispatcher(ctx *context.ApplicationContext) *Dispatcher {
	return &Dispatcher{
		ctx: ctx,
		resolvers: []ArgumentResolver{
			&EchoContextResolver{},
			&PathVariableResolver{},
			&RequestParamResolver{},
		},
		returnHandlers: []ReturnValueHandler{
			&ErrorReturnValueHandler{},
			&StringReturnValueHandler{},
			&JSONReturnValueHandler{},
		},
	}
}

func (d *Dispatcher) RegisterRoutes(e *echo.Echo) {
	d.handlers = CreateControllerMappings(d.ctx)

	e.Any("/*", func(c echo.Context) error {
		return d.dispatch(c)
	})
}

func (d *Dispatcher) dispatch(c echo.Context) error {
	requestPath := c.Request().URL.Path
	requestMethod := c.Request().Method

	for _, h := range d.handlers {
		// HTTP Method 먼저 필터링
		if h.Mapping.Method != requestMethod {
			continue
		}

		// Path 매칭
		vars, ok := matchPatch(h.Mapping.Path, requestPath)
		if !ok {
			continue
		}

		// 매칭 성공 -> 실행
		rc := &RequestContext{
			Echo:     c,
			PathVars: vars,
		}

		return d.invokeHandlerMethod(h, rc)
	}

	// 아무 핸들러도 못 찾음
	return echo.ErrNotFound
}

func (d *Dispatcher) invokeHandlerMethod(h HandlerMethod, rc *RequestContext) error {
	methodType := h.Method.Type
	args := []reflect.Value{
		reflect.ValueOf(h.Bean),
	}

	// Argument Resolver
	for i := 1; i < methodType.NumIn(); i++ {
		paramType := methodType.In(i)

		value, err := d.resolveArgument(rc, paramType, i-1)
		if err != nil {
			return err
		}

		args = append(args, value)
	}

	// 메서드 호출
	results := h.Method.Func.Call(args)

	// 반환값 처리
	if len(results) == 0 {
		return nil
	}

	return d.handleReturnValue(results[0], rc)
}

func (d *Dispatcher) handleReturnValue(value reflect.Value, rc *RequestContext) error {
	for _, handler := range d.returnHandlers {
		if handler.Supports(value.Type()) {
			return handler.Handle(value, rc.Echo)
		}
	}

	return fmt.Errorf("%v 반환 값을 처리할 핸들러가 없습니다.", value.Type())
}

func (d *Dispatcher) resolveArgument(rc *RequestContext, t reflect.Type, paramIndex int) (reflect.Value, error) {
	for _, resolver := range d.resolvers {
		if !resolver.Supports(t) {
			continue
		}

		value, err := resolver.Resolve(rc, t, paramIndex)
		if err == nil {
			return value, nil
		}
	}
	return reflect.Value{}, fmt.Errorf("리졸버 %v에 대해 인자가 없습니다.", t)
}
