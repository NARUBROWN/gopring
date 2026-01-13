package context

import "reflect"

// 빈 메타데이터
type BeanDefinition struct {
	BeanType reflect.Type
	Factory  func(ctx *ApplicationContext) any
}
