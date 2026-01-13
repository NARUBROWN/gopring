package example

import (
	"fmt"
	"gopring/web"
)

type UserController struct {
	Service *UserService
}

func (u *UserController) Mappings() map[string]web.RequestMapping {
	return map[string]web.RequestMapping{
		"Users": {
			Method: "GET",
			Path:   "/users",
		},
		"GetUser": {
			Method: "GET",
			Path:   "/users/{id}",
		},
	}
}

func (u *UserController) Users(name string) string {
	return fmt.Sprintf("name=%s", name)
}

func (u *UserController) GetUser(id int) string {
	return fmt.Sprintf("id=%d", id)
}
