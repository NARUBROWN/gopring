package context

import "reflect"

// 핵심 컨테이너 (Application Context)
type ApplicationContext struct {
	BeanDefinitions map[reflect.Type]BeanDefinition
	SingletonBeans  map[reflect.Type]any
	CreatingBeans   map[reflect.Type]bool
}

func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{
		BeanDefinitions: map[reflect.Type]BeanDefinition{},
		SingletonBeans:  map[reflect.Type]any{},
		CreatingBeans:   map[reflect.Type]bool{},
	}
}
