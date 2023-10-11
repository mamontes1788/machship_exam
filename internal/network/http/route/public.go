package route

import (
	"machship/internal/network/http/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	usersHandler handler.UsersHandler
}

func NewGenericRoute(usersHandler handler.UsersHandler) Route {
	route := new(Route)
	route.usersHandler = usersHandler
	return *route
}

func (c *Route) PublicSetup(routeGroup *gin.RouterGroup) {
	publicGroup := routeGroup.Group("v1/")
	publicGroup.GET("/retrieveUsers", c.usersHandler.GetUsers)
}
