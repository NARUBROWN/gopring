package app

import (
	"gopring/context"
	"gopring/web"

	"github.com/labstack/echo/v4"
)

func Bootstrap() *echo.Echo {
	// ApplicationContext 생성
	applicationContext := context.NewApplicationContext()

	// BeanDefinition 등록
	registerBean(applicationContext)

	// Dispatcher 생성 (DispatcherServlet)
	dispatcher := web.NewDispatcher(applicationContext)

	// Container 생성 (Servlet Container)
	e := echo.New()

	// DispatcherServlet.onStartup()
	dispatcher.RegisterRoutes(e)

	return e
}
