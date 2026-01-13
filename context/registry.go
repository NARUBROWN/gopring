package context

import "reflect"

func (ctx *ApplicationContext) RegisterBean(beanType reflect.Type, factory func(*ApplicationContext) any) {
	ctx.BeanDefinitions[beanType] = BeanDefinition{
		BeanType: beanType,
		Factory:  factory,
	}
}

func (ctx *ApplicationContext) GetBean(beanType reflect.Type) any {
	// BeanType이 일치하면, Application Context에서 Bean 반환
	if bean, ok := ctx.SingletonBeans[beanType]; ok {
		return bean
	}

	// Bean이 생성중이면 오류!
	if ctx.CreatingBeans[beanType] {
		panic("순환 의존성 오류: " + beanType.String())
	}

	definition, ok := ctx.BeanDefinitions[beanType]
	// Bean 정의가 없으면 오류
	if !ok {
		panic("빈 정의가 없습니다: " + beanType.String())
	}

	// Application Context에 Bean Type을 Key로 Bean이 생성 중이라고 표시
	ctx.CreatingBeans[beanType] = true
	bean := definition.Factory(ctx)
	// Application Context에 Bean Type을 Key로 Bean 생성이 끝났다고 표시
	ctx.CreatingBeans[beanType] = false

	// Application Context에 Bean 최종 등록
	ctx.SingletonBeans[beanType] = bean
	return bean
}
