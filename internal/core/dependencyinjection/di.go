package dependencyinjection

import (
	"machship/internal/client"
	"machship/internal/core/service"
	"machship/internal/network/http"
	"machship/internal/network/http/handler"
	"machship/internal/network/http/route"
)

func InitializeAPIs() *http.ServerHTTP {
	apiClient := client.NewApiRequest()
	userService := service.NewUserServiceImpl(apiClient)
	usersHandler := handler.NewUsersHandler(userService)
	router := route.NewGenericRoute(usersHandler)
	serverHTTP := http.NewServerHTTP(router)
	return serverHTTP
}
