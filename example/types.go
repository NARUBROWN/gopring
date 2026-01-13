package example

import "reflect"

func UserRepositoryType() reflect.Type {
	return reflect.TypeOf(&UserRepository{})
}

func UserServiceType() reflect.Type {
	return reflect.TypeOf(&UserService{})
}

func UserControllerType() reflect.Type {
	return reflect.TypeOf(&UserController{})
}
