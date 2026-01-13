package app

import (
	"gopring/context"
	"gopring/example"
)

func registerBean(ctx *context.ApplicationContext) {
	// Repository
	ctx.RegisterBean(
		example.UserRepositoryType(),
		func(ac *context.ApplicationContext) any {
			return &example.UserRepository{}
		},
	)

	// Service
	ctx.RegisterBean(
		example.UserServiceType(),
		func(ac *context.ApplicationContext) any {
			return &example.UserService{
				Repository: ac.GetBean(
					example.UserRepositoryType(),
				).(*example.UserRepository),
			}
		},
	)

	// Controller
	ctx.RegisterBean(
		example.UserControllerType(),
		func(ac *context.ApplicationContext) any {
			return &example.UserController{
				Service: ac.GetBean(
					example.UserServiceType(),
				).(*example.UserService),
			}
		},
	)
}
